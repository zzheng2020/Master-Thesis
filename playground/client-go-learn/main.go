package main

import (
	"context"
	"database/sql"
	"fmt"
	"path/filepath"

	_ "github.com/lib/pq"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	// create a Kubernetes clientset
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
			panic(err.Error())
		}
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// get the IP address of the postgres-1 Pod
	podName := "pgupgrade-sample-689bfd9c4-xwlqk"
	namespace := "default"
	pod, err := clientset.CoreV1().Pods(namespace).Get(context.Background(), podName, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}
	podIP := pod.Status.PodIP
	fmt.Println(podIP)

	// connect to the PostgreSQL database
	connectionString := fmt.Sprintf("host=%s port=56001 user=postgres password=mysecretpassword dbname=mydatabase sslmode=disable", "127.0.0.1")
	fmt.Println(connectionString)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// send a request to the table
	insertStatement := "insert into table_1 values(1000000)"
	// createPub := "create publication pub2 for all tables"
	// deletePub := "delete from pg_catalog.pg_publication where pubname='pub2'"
	_, err = db.Exec(insertStatement)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Statement executed successfully")
}
