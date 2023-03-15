package main

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"path/filepath"
	"time"

	_ "github.com/lib/pq"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	// Load Kubernetes configuration
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

	// Create a Kubernetes clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// Get the PostgreSQL service
	serviceName := "postgres-master"
	service, err := clientset.CoreV1().Services("default").Get(context.Background(), serviceName, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("service: ", service.Spec.ClusterIP)

	// Create a connection string to the PostgreSQL database
	// ! NOTE: we need to run `minikube service <service-name>` to forward the port.
	targetPort := 60227
	connString := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		"127.0.0.1", targetPort, "mydatabase", "postgres", "mysecretpassword")

	// Create a ticker for calling read() every 10 seconds
	readTicker := time.NewTicker(10 * time.Second)
	defer readTicker.Stop()

	// Create a ticker for calling write() every 30 seconds
	writeTicker := time.NewTicker(30 * time.Second)
	defer writeTicker.Stop()

	for {
		select {
		case <-readTicker.C:
			read(connString)
		case <-writeTicker.C:
			write(connString)
		}
	}

}

func read(connString string) {
	fmt.Println("====== begin the read ======")

	// Connect to the PostgreSQL database
	db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Send a read query to the database
	rows, err := db.Query("SELECT * FROM table_1")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// Print the results of the read query
	for rows.Next() {
		var id string
		err := rows.Scan(&id)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("id: %s\n", id)
	}

	fmt.Println("====== finish the read ======")
}

func write(connString string) {
	// Connect to the PostgreSQL database
	db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	for {
		randomID := rand.Intn(1000)
		fmt.Println("randomID: ", randomID)
		// Send a write query to the database
		_, err = db.Exec("INSERT INTO table_1 VALUES ($1)", randomID)
		if err != nil {
			panic(err.Error())
		} else {
			break
		}
	}

	fmt.Println("====== suceessfully write to database ======")
}
