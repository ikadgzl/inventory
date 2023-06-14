package db

import (
	"database/sql"
	"log"
	"time"

	"github.com/ikadgzl/inventory/authsvc/internal/config"
	_ "github.com/lib/pq"
)

func NewPostgresDB(c *config.DBConfig) *sql.DB {
	db, err := sql.Open("postgres", c.ConnectionString())
	if err != nil {
		log.Fatalf("failed to open db connection: %v", err)
	}

	for i := 0; i < 5; i++ {
		err = db.Ping()
		if err == nil {
			break
		}

		log.Printf("trying to ping db, retry: %d", i+1)
		time.Sleep(1 * time.Second)

		if i == 4 {
			log.Fatalf("failed to ping db: %v", err)
		}
	}

	c.Configure(db)

	return db
}
