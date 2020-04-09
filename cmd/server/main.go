package main

import (
	"log"

	bcserver "github.com/stoina/go-blockchain/cmd/server/bcserver"
)

func main() {
	bcserver := bcserver.New()
	err := bcserver.Start()

	if err != nil {
		log.Println(err.Error())
	}

	log.Println("Server successfully started...")
}
