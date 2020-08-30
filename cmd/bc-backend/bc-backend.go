package main

import (
	"fmt"
	"net/http"
	"strconv"

	controller "github.com/Stoina/go-blockchain/cmd/bc-backend/controller"
	db "github.com/Stoina/go-database"
)

func main() {
	port := 8080
	dbConn, err := db.OpenDBConnection("postgres", "localhost", 5432, "blockchain_admin", "Kfpqvvc2BODNGxPZkZeY", "blockchain")

	if err != nil {
		fmt.Println(err.Error())
	}

	controllers := getControllerToHandle(dbConn)

	for _, controller := range controllers {
		http.Handle("/"+controller.URL()+"/", controller)
	}

	fmt.Println(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}

func getControllerToHandle(dbConn *db.Connection) []controller.Controller {
	return []controller.Controller{
		controller.NewParticipantController(dbConn),
		controller.NewTransactionController(dbConn),
		controller.NewBlockController(dbConn)}
}
