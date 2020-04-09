package bcserver

import (
	"net/http"
	"strconv"

	controller "github.com/stoina/go-blockchain/cmd/server/bcserver/controller"
	db "github.com/stoina/go-database"
)

// BCServer exported
// ...
type BCServer struct {
	Port   int
	DBConn *db.Connection
}

// New exported
// ...
func New() *BCServer {
	return &BCServer{
		Port: 8080}
}

// Start exported
// ...
func (server *BCServer) Start() error {
	dbConn, err := db.OpenDBConnection("postgres", "localhost", 5432, "blockchain_admin", "Kfpqvvc2BODNGxPZkZeY", "blockchain")

	if err != nil {
		return err
	}

	server.DBConn = dbConn

	controllers := getControllerToHandle(dbConn)

	for _, controller := range controllers {
		http.Handle("/"+controller.URL()+"/", controller)
	}

	return http.ListenAndServe(":"+strconv.Itoa(server.Port), nil)
}

func getControllerToHandle(dbConn *db.Connection) []controller.Controller {
	return []controller.Controller{
		controller.NewParticipantController(dbConn),
		controller.NewTransactionController(dbConn),
		controller.NewBlockController(dbConn)}
}
