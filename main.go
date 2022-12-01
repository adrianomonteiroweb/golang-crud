package main

import (
	"database/sql"
	"log"

	"github.com/adrianomonteiroweb/golang-crud/api"
	db "github.com/adrianomonteiroweb/golang-crud/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://postgres:postgres@localhost:5438/products_db?sslmode=disable"
	serverAddress = "0.0.0.0:3333"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot connect db:", err)
	}

	store := db.NewStoreExec(conn)
	server := api.InstanceServer(store)

	err = server.Start(serverAddress)

	if err != nil {
		log.Fatal("api started with error:", err)
	}
}