package db

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v4"
)

const (
	maxRetries    = 5
	retryInterval = 5 * time.Second
)

var Conn *pgx.Conn

func Init(connStr string) {
	var err error

	for i := 0; i < maxRetries; i++ {
		Conn, err = pgx.Connect(context.Background(), connStr)
		if err == nil {
			break
		}
		log.Printf("Failed to connect to database: %v. Retrying in %v...", err, retryInterval)

		time.Sleep(retryInterval)
	}

	if err != nil {
		log.Fatalf("Failed to connect to database")
	}

	log.Printf("Successful connect to database")

}
