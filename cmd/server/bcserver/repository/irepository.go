package repository

import (
	"github.com/stoina/go-blockchain/cmd/server/bcserver/repository/result"
	db "github.com/stoina/go-database"
)

// Repository exported
// Repository ...
type Repository interface {
	Create(jsonString string, db *db.Connection) *result.RepositoryResult
	Update(jsonString string, db *db.Connection) *result.RepositoryResult
	Delete(jsonString string, db *db.Connection) *result.RepositoryResult
	ReadAll(db *db.Connection) *result.RepositoryResult
	ReadByID(id string, db *db.Connection) *result.RepositoryResult
	ReadByProperty(propertyName string, value interface{}, db *db.Connection) *result.RepositoryResult
}
