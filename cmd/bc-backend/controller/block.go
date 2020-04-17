package controller

import (
	"net/http"
	"net/url"

	repo "github.com/stoina/go-blockchain/cmd/bc-backend/repository"
	result "github.com/stoina/go-blockchain/cmd/bc-backend/repository/result"
	db "github.com/stoina/go-database"
)

// BlockController exported
// ...
type BlockController struct {
	url        string
	name       string
	repository repo.Repository
	dbConn     *db.Connection
}

// NewBlockController exorted
// ...
func NewBlockController(dbConn *db.Connection) *BlockController {
	return &BlockController{
		name:       "Block Controller",
		url:        "blocks",
		repository: repo.NewBlockRepository(),
		dbConn:     dbConn}
}

// URL exported
// ...
func (bc *BlockController) URL() string {
	return bc.url
}

// Name exported
// ...
func (bc *BlockController) Name() string {
	return bc.name
}

// Repository exported
// ...
func (bc *BlockController) Repository() repo.Repository {
	return bc.repository
}

// ServeHTTP exported
// ...
func (bc *BlockController) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	HandleRequest(writer, request, bc)
}

// Post exported
// ...
func (bc *BlockController) Post(contentType string, content string) *result.RepositoryResult {
	return nil
}

// Put exported
// ...
func (bc *BlockController) Put(par string) *result.RepositoryResult {
	return nil
}

// Patch exported
// ...
func (bc *BlockController) Patch(par string) *result.RepositoryResult {
	return nil
}

// Delete exported
// ...
func (bc *BlockController) Delete(par string) *result.RepositoryResult {
	return nil
}

// Get exported
// ...
func (bc *BlockController) Get(calledURL *url.URL) *result.RepositoryResult {
	return nil
}
