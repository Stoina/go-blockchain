package controller

import (
	"net/http"
	"net/url"

	repo "github.com/stoina/go-blockchain/cmd/server/bcserver/repository"
	result "github.com/stoina/go-blockchain/cmd/server/bcserver/repository/result"
)

// Controller exported
// ...
type Controller interface {
	URL() string
	Name() string
	Repository() repo.Repository

	ServeHTTP(writer http.ResponseWriter, request *http.Request)
	Post(contentType string, content string) *result.RepositoryResult
	Put(par string) *result.RepositoryResult
	Patch(par string) *result.RepositoryResult
	Delete(par string) *result.RepositoryResult
	Get(calledURL *url.URL) *result.RepositoryResult
}
