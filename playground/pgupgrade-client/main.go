package main

import (
	"context"
	"database/sql"
	"fmt"
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
	// serviceName := "postgres-master"
	serviceName := "nginx-service"
	service, err := clientset.CoreV1().Services("default").Get(context.Background(), serviceName, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("service: ", service.Spec.ClusterIP)

	// Create a connection string to the PostgreSQL database
	// ! NOTE: we need to run `minikube service <service-name>` to forward the port.
	targetPort := 54165
	connString := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		"127.0.0.1", targetPort, "mydatabase", "postgres", "mysecretpassword")

	// Create a ticker for calling read() every 10 seconds
	readTicker := time.NewTicker(5 * time.Second)
	defer readTicker.Stop()

	// Create a ticker for calling write() every 30 seconds
	writeTicker := time.NewTicker(10 * time.Second)
	defer writeTicker.Stop()

	// read from database every 10 seconds
	// for {
	// 	read(connString)
	// 	time.Sleep(10 * time.Second)
	// }

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
		// panic(err.Error())
		fmt.Println(err.Error())
	}
	defer db.Close()

	// Send a read query to the database
	query := "SELECT * FROM (SELECT id FROM table_1 ORDER BY id DESC LIMIT 5) AS top_five ORDER BY id DESC;"
	rows, err := db.Query(query)
	if err != nil {
		// panic(err.Error())
		fmt.Println(err.Error())
	}
	defer rows.Close()

	// Print the results of the read query
	for rows.Next() {
		var id string
		err := rows.Scan(&id)
		if err != nil {
			// panic(err.Error())
			fmt.Println(err.Error())

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

	startNum := 1000001
	for {
		// Send a write query to the database
		_, err = db.Exec("INSERT INTO table_1 VALUES ($1)", startNum)
		if err != nil {
			startNum++
		} else {
			break
		}
	}

	fmt.Println("====== suceessfully write to database ======")
}
