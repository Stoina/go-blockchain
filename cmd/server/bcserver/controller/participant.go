package controller

import (
	"net/http"
	"net/url"

	db "github.com/stoina/go-database"

	repo "github.com/stoina/go-blockchain/cmd/server/bcserver/repository"
	result "github.com/stoina/go-blockchain/cmd/server/bcserver/repository/result"
)

// ParticipantController exported
// ...
type ParticipantController struct {
	url        string
	name       string
	repository repo.Repository
	dbConn     *db.Connection
}

// NewParticipantController exorted
// ...
func NewParticipantController(dbConn *db.Connection) *ParticipantController {
	return &ParticipantController{
		name:       "Participant Controller",
		url:        "participant",
		repository: repo.NewParticipantRepository(),
		dbConn:     dbConn}
}

// URL exported
// ...
func (pc *ParticipantController) URL() string {
	return pc.url
}

// Name exported
// ...
func (pc *ParticipantController) Name() string {
	return pc.name
}

// Repository exported
// ...
func (pc *ParticipantController) Repository() repo.Repository {
	return pc.repository
}

func (pc *ParticipantController) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	HandleRequest(writer, request, pc)
}

// Post exported
// ...
func (pc *ParticipantController) Post(contentType string, content string) *result.RepositoryResult {
	return pc.repository.Create(content, pc.dbConn)
}

// Put exported
// ...
func (pc *ParticipantController) Put(par string) *result.RepositoryResult {
	return nil

}

// Patch exported
// ...
func (pc *ParticipantController) Patch(par string) *result.RepositoryResult {
	return nil

}

// Delete exported
// ...
func (pc *ParticipantController) Delete(par string) *result.RepositoryResult {
	return nil

}

// Get exported
// ...
func (pc *ParticipantController) Get(calledURL *url.URL) *result.RepositoryResult {
	return pc.repository.ReadAll(pc.dbConn)
}
