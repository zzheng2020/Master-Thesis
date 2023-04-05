package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 60227
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

	// for range to exec read func
	for range readTicker.C {
		read(db)
	}
}

func read(db *sql.DB) {
	fmt.Println("====== begin the read ======")

	// readQuery := `SELECT * FROM table_1`
	readQuery := `SELECT application_name, client_addr, state, sync_state, pg_size_pretty(pg_wal_lsn_diff(pg_current_wal_flush_lsn(), sent_lsn)) AS replication_lag FROM pg_stat_replication;`
	rows, err := db.Query(readQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Records in the database:")
	for rows.Next() {
		var applicationName, clientAddr, state, syncState, replicationLag string
		err := rows.Scan(&applicationName, &clientAddr, &state, &syncState, &replicationLag)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("Application Name: %s, Client Addr: %s, State: %s, Sync State: %s, Replication Lag: %s\n", applicationName, clientAddr, state, syncState, replicationLag)
	}

	fmt.Println("====== finish the read ======")
}
