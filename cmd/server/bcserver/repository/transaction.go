package repository

import (
	"github.com/stoina/go-blockchain/cmd/server/bcserver/repository/result"
	db "github.com/stoina/go-database"
)

// TransactionRepository exported
// ...
type TransactionRepository struct {
}

// NewTransactionRepository exported
// ...
func NewTransactionRepository() *TransactionRepository {
	return &TransactionRepository{}
}

// Create exported
// ...
func (tp *TransactionRepository) Create(jsonString string, db *db.Connection) *result.RepositoryResult {
	return nil
}

// Update exported
// ...
func (tp *TransactionRepository) Update(jsonString string, db *db.Connection) *result.RepositoryResult {
	return nil
}

// Delete exported
// ...
func (tp *TransactionRepository) Delete(jsonString string, db *db.Connection) *result.RepositoryResult {
	return nil
}

// ReadAll exported
// ...
func (tp *TransactionRepository) ReadAll(db *db.Connection) *result.RepositoryResult {
	return nil
}

// ReadByID exported
// ...
func (tp *TransactionRepository) ReadByID(id string, db *db.Connection) *result.RepositoryResult {
	return nil
}

// ReadByProperty exported
// ...
func (tp *TransactionRepository) ReadByProperty(propertyName string, value interface{}, db *db.Connection) *result.RepositoryResult {
	return nil
}
