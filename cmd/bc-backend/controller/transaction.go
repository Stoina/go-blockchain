package controller

import (
	"net/http"
	"net/url"

	db "github.com/stoina/go-database"

	repo "github.com/stoina/go-blockchain/cmd/bc-backend/repository"
	result "github.com/stoina/go-blockchain/cmd/bc-backend/repository/result"
)

// TransactionController exported
// ...
type TransactionController struct {
	url        string
	name       string
	repository repo.Repository
	dbConn     *db.Connection
}

// NewTransactionController exorted
// ...
func NewTransactionController(dbConn *db.Connection) *TransactionController {
	return &TransactionController{
		name:       "Transaction Controller",
		url:        "transactions",
		repository: repo.NewTransactionRepository(),
		dbConn:     dbConn}
}

// URL exported
// ...
func (tc *TransactionController) URL() string {
	return tc.url
}

// Name exported
// ...
func (tc *TransactionController) Name() string {
	return tc.name
}

// Repository exported
// ...
func (tc *TransactionController) Repository() repo.Repository {
	return tc.repository
}

func (tc *TransactionController) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	HandleRequest(writer, request, tc)
}

// Post exported
// ...
func (tc *TransactionController) Post(contentType string, content string) *result.RepositoryResult {
	return tc.repository.Create(content, tc.dbConn)
}

// Put exported
// ...
func (tc *TransactionController) Put(par string) *result.RepositoryResult {
	return nil

}

// Patch exported
// ...
func (tc *TransactionController) Patch(par string) *result.RepositoryResult {
	return nil

}

// Delete exported
// ...
func (tc *TransactionController) Delete(par string) *result.RepositoryResult {
	return nil

}

// Get exported
// ...
func (tc *TransactionController) Get(calledURL *url.URL) *result.RepositoryResult {
	return nil
}
