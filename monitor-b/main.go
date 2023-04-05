package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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
	r := gin.Default()

	replicationStatusChannel := make(chan []map[string]string)
	go startReplicationStatusTicker(replicationStatusChannel)

	r.GET("/sync", func(c *gin.Context) {
		replicationStatus := <-replicationStatusChannel
		c.JSON(http.StatusOK, replicationStatus)
	})

	r.Run(":8080") // Listen and serve on 0.0.0.0:8080
}

func startReplicationStatusTicker(replicationStatusChannel chan []map[string]string) {
	ticker := time.NewTicker(5 * time.Second) // Adjust the duration to your desired update interval

	for range ticker.C {
		replicationStatus := getReplicationStatus()
		replicationStatusChannel <- replicationStatus
	}
}

func getReplicationStatus() []map[string]string {
	// Setup PostgreSQL connection here
	// connStr := "user=username dbname=database password=password host=hostname port=port sslmode=disable"
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	readQuery := `SELECT application_name, client_addr, state, sync_state, pg_size_pretty(pg_wal_lsn_diff(pg_current_wal_flush_lsn(), sent_lsn)) AS replication_lag FROM pg_stat_replication;`
	rows, err := db.Query(readQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var replicationStatus []map[string]string
	for rows.Next() {
		var applicationName, clientAddr, state, syncState, replicationLag string
		err := rows.Scan(&applicationName, &clientAddr, &state, &syncState, &replicationLag)
		if err != nil {
			panic(err.Error())
		}

		record := map[string]string{
			"application_name": applicationName,
			"client_addr":      clientAddr,
			"state":            state,
			"sync_state":       syncState,
			"replication_lag":  replicationLag,
		}
		replicationStatus = append(replicationStatus, record)
		fmt.Println(record)
	}

	return replicationStatus
}
