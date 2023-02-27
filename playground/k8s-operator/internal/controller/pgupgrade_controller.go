/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"bytes"
	"context"
	"fmt"
	"path/filepath"
	"time"

	"github.com/go-logr/logr"

	"sigs.k8s.io/controller-runtime/pkg/log"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/remotecommand"
	"k8s.io/client-go/util/homedir"
	"k8s.io/kubectl/pkg/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"k8s.io/apimachinery/pkg/api/errors"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	pgupgradev1 "zzh.domain/pgoperator/api/v1"

	_ "github.com/lib/pq"
)

// PgUpgradeReconciler reconciles a PgUpgrade object
type PgUpgradeReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	Log    logr.Logger
}

//+kubebuilder:rbac:groups=pgupgrade.zzh.domain,resources=pgupgrades,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=pgupgrade.zzh.domain,resources=pgupgrades/status,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=pgupgrade.zzh.domain,resources=pgupgrades/finalizers,verbs=get;list;watch;create;update;patch;delete

//+kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=apps,resources=services,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups="",resources=services,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups="",resources=pods,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups="",resources=pods/exec,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the PgUpgrade object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *PgUpgradeReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	// _ = log.FromContext(ctx)

	// log := r.Log.WithValues("PgUpgrade", req.NamespacedName)
	log := log.FromContext(ctx)

	log.V(1).Info("Reconcile function starts.")

	// TODO(user): your logic here

	// Get the PgUpgrade instance, and check if the instance exists.
	instance := &pgupgradev1.PgUpgrade{}
	err := r.Get(ctx, req.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			log.V(1).Info("PgUpgrade resource not found.")
			return ctrl.Result{}, nil
		}

		log.Error(err, "Failed to get PgUpgrade.")
		return ctrl.Result{RequeueAfter: time.Second * 5}, err
	}

	// Check if the PgUpgrade deployment already exists, and if NOT, create a new one.
	found := &appsv1.Deployment{}
	err = r.Get(ctx, types.NamespacedName{
		Name:      instance.Name,
		Namespace: instance.Namespace,
	}, found)
	if err != nil && errors.IsNotFound(err) {
		dep := r.deploymentForPgUpgrade(instance)
		log.V(1).Info("Creating a new Deployment", "Deployment.Namespace", dep.Namespace, "Deployment.Name", dep.Name)
		// Create a new Deployment in K8s cluster.
		err = r.Create(ctx, dep)
		if err != nil {
			log.Error(err, "Failed to create a new Deployment.", "Deployment.Namespace", dep.Namespace, "Deployment.Name", dep.Name)
			return ctrl.Result{RequeueAfter: time.Second * 5}, err
		}
		return ctrl.Result{Requeue: true}, nil
	} else if err != nil {
		log.Error(err, "Failed to get Deployment")
		return ctrl.Result{RequeueAfter: time.Second * 5}, err
	}

	// Ensure the deployment image is the same as the spec.
	image := instance.Spec.Image
	needUpdate := false
	if (*found).Spec.Template.Spec.Containers[0].Image != image {
		log.V(1).Info("Deployment spec.template.spec.container[0].image change", "from", (*found).Spec.Template.Spec.Containers[0].Image, "to", image)
		found.Spec.Template.Spec.Containers[0].Image = image
		needUpdate = true
	}
	// If the deployment image is not the same as the spec, update the deployment.
	if needUpdate {
		err = r.Update(ctx, found)
		if err != nil {
			log.Error(err, "Failed to update Deployment.", "Deployment.Namespace", found.Namespace, "Deployment.Name", found.Name)
			return ctrl.Result{RequeueAfter: time.Second * 5}, err
		}
		return ctrl.Result{Requeue: true}, nil
	}

	// Check if the PgUpgrade service already exists, and if NOT, create a new one.
	foundService := &corev1.Service{}
	err = r.Get(ctx, types.NamespacedName{Name: instance.Name + "-service", Namespace: instance.Namespace}, foundService)
	if err != nil && errors.IsNotFound(err) {
		// Define a new service
		svc := r.serviceForPgUpgrade(instance)
		log.V(1).Info("Creating a new Service", "Service.Namespace", svc.Namespace, "Service.Name", svc.Name)
		err = r.Create(ctx, svc)
		if err != nil {
			log.Error(err, "Failed to create a new Service.", "Service.Namespace", svc.Namespace, "Service.Name", svc.Name)
			return ctrl.Result{RequeueAfter: time.Second * 5}, err
		}
		return ctrl.Result{Requeue: true}, nil
	} else if err != nil {
		log.Error(err, "Failed to get Service")
		return ctrl.Result{RequeueAfter: time.Second * 5}, err
	}

	// If old db has already been dumped and it didn't update before, then start to create the subscription.
	if instance.Spec.PgDump && !instance.Status.Upgrade {
		log.V(1).Info("Start to create subscriptions.")
		instance.Status.Upgrade = true
		err = r.createSubscriptions(ctx, instance, *found)
		if err != nil {
			log.Error(err, "Failed to create subscriptions.")
			return ctrl.Result{RequeueAfter: time.Second * 5}, err
		}
		log.V(1).Info("Successfully create subscriptions.")
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *PgUpgradeReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&pgupgradev1.PgUpgrade{}).
		Complete(r)
}

// create a new Deployment for a PgUpgrade resource.
func (r *PgUpgradeReconciler) deploymentForPgUpgrade(pg *pgupgradev1.PgUpgrade) *appsv1.Deployment {
	labels := labelsForPgUpgrade(pg.Name)
	dep := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      pg.Name,
			Namespace: pg.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Image: pg.Spec.Image,
						Name:  "pgupgrade",
						Env: []corev1.EnvVar{
							{
								Name:  "POSTGRES_USER",
								Value: "postgres",
							},
							{
								Name:  "POSTGRES_PASSWORD",
								Value: "mysecretpassword",
							},
							{
								Name:  "POSTGRES_DB",
								Value: "mydatabase",
							},
						},
					}},
				},
			},
		},
	}
	// Set PgUpgrade instance as the owner and controller.
	ctrl.SetControllerReference(pg, dep, r.Scheme)

	return dep
}

// create a new Service for a PgUpgrade resource.
func (r *PgUpgradeReconciler) serviceForPgUpgrade(pg *pgupgradev1.PgUpgrade) *corev1.Service {
	labels := labelsForPgUpgrade(pg.Name)
	svc := &corev1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      pg.Name + "-service",
			Namespace: pg.Namespace,
			Labels:    labels,
		},
		Spec: corev1.ServiceSpec{
			Type: corev1.ServiceTypeNodePort,
			Ports: []corev1.ServicePort{{
				Port: 5432,
				TargetPort: intstr.IntOrString{
					IntVal: 5432,
				},
			}},
			Selector: labels,
		},
	}
	// Set PgUpgrade instance as the owner and controller.
	ctrl.SetControllerReference(pg, svc, r.Scheme)

	return svc
}

func labelsForPgUpgrade(name string) map[string]string {
	return map[string]string{"app": "pgupgrade", "pgupgrade_cr": name}
}

// Create subscriptions.
func (r *PgUpgradeReconciler) createSubscriptions(ctx context.Context, pg *pgupgradev1.PgUpgrade, found appsv1.Deployment) error {
	log := log.FromContext(ctx)
	// Create a new clientset
	var kubeconfig string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = filepath.Join(home, ".kube", "config")
	} else {
		kubeconfig = ""
	}
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		config, err = rest.InClusterConfig()
		if err != nil {
			return err
		}
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	// Get the pod name of Deployment.
	pods, err := clientset.CoreV1().Pods("default").List(ctx, metav1.ListOptions{
		LabelSelector: fmt.Sprintf("app=%s", "pgupgrade"),
	})
	if err != nil {
		return err
	}
	podName := pods.Items[0].Name
	namespace := "default"
	pod, err := clientset.CoreV1().Pods(namespace).Get(context.Background(), podName, metav1.GetOptions{})
	if err != nil {
		return err
	}
	podIP := pod.Status.PodIP

	log.V(1).Info("Pod IP and Name", "Pod Ip is ", podIP, "Pod Name is ", podName)
	// log.V(1).Info("Pod IP and Name", "Pod Ip is ", podIP, "Pod Name is ", podName)

	// Execute the command in the pod.
	subscription := fmt.Sprintf("create subscription %s connection 'dbname=%s host=%s user=postgres password=postgres port=%s' publication %s;", pg.Spec.SubName, pg.Spec.DBName, pg.Spec.OldDBHost, pg.Spec.OldDBPort, pg.Spec.PubName)
	// cmd := []string{"psql", "-U", "postgres", "mydatabase", "-c", subscription}
	cmd := []string{"psql", "-U", "postgres", pg.Spec.DBName, "-c", subscription}

	req := clientset.CoreV1().RESTClient().
		Post().
		Namespace(namespace).
		Resource("pods").
		Name(podName).
		SubResource("exec").
		VersionedParams(&corev1.PodExecOptions{
			Command: cmd,
			Stdin:   false,
			Stdout:  true,
			Stderr:  true,
		}, scheme.ParameterCodec)

	exec, err := remotecommand.NewSPDYExecutor(config, "POST", req.URL())
	if err != nil {
		return err
	}
	// Get the output of the command.
	var stdout, stderr bytes.Buffer
	err = exec.StreamWithContext(ctx, remotecommand.StreamOptions{
		Stdin:  nil,
		Stdout: &stdout,
		Stderr: &stderr,
		Tty:    false,
	})
	if err != nil {
		return err
	}
	fmt.Println(stdout.String())

	return nil
}
