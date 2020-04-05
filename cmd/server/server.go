package main

import (
	"log"

	db "github.com/Stoina/go-database"
)

func main() {

	_, err := initializeDBConnection()

	if err != nil {
		log.Println(err.Error())
	}
}

func initializeDBConnection() (*db.Connection, error) {
	return db.OpenDBConnection("postgres", "localhost", 5432, "blockchain_admin", "Kfpqvvc2BODNGxPZkZeY", "blockchain")
}
