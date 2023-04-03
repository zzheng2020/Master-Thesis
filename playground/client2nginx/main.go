package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 80
	// port     = 60227
	user     = "postgres"
	password = "mysecretpassword"
	dbname   = "mydatabase"
)

func main() {
	// Connect to the PostgreSQL database through Nginx
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Connected to the database!")

	readTicker := time.NewTicker(10 * time.Second)
	defer readTicker.Stop()

	// Create a ticker for calling write() every 30 seconds
	writeTicker := time.NewTicker(30 * time.Second)
	defer writeTicker.Stop()

	for {
		select {
		case <-readTicker.C:
			read(db)
		case <-writeTicker.C:
			write(db)
		}
	}
}

func read(db *sql.DB) {
	fmt.Println("====== begin the read ======")

	readQuery := `SELECT * FROM table_1`
	rows, err := db.Query(readQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Records in the database:")
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

func write(db *sql.DB) {
	// Connect to the PostgreSQL database

	for {
		randomID := rand.Intn(1000)
		fmt.Println("randomID: ", randomID)
		// Send a write query to the database
		_, err := db.Exec("INSERT INTO table_1 VALUES ($1)", randomID)
		if err != nil {
			continue
		} else {
			break
		}
	}

	fmt.Println("====== suceessfully write to database ======")
}
