package main

import (
	"log"

	bcserver "github.com/stoina/go-blockchain/cmd/server/bcserver"
)

func main() {
	bcserver := bcserver.Create()
	err := bcserver.Start()

	if err != nil {
		log.Println("Server successfully started...")
	}
}
