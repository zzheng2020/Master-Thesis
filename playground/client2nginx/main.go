package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 52993
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

	// Read records from the database
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
}
