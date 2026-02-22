package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func Connect() *sql.DB {
	user := os.Getenv("PG_USER")
	pass := os.Getenv("PG_PASS")
	dbName := os.Getenv("PG_DB")
	host := os.Getenv("PG_HOST")
	port := os.Getenv("PG_PORT")
	
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, pass, host, port, dbName)
	db, err := sql.Open("pgx", dsn) // pgx is the db driver
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to PostgreSQL")
	return db
}