package repository

import (
	"encoding/json"

	"github.com/stoina/go-blockchain/blockchain/transaction"
	transactiondb "github.com/stoina/go-blockchain/blockchain/transaction/db"
	"github.com/stoina/go-blockchain/blockchain/util/encoding"
	"github.com/stoina/go-blockchain/blockchain/uuid"
	"github.com/stoina/go-blockchain/cmd/bc-backend/repository/result"
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

	var transaction transaction.Transaction
	err := json.Unmarshal([]byte(jsonString), &transaction)

	if err != nil {
		return result.New("", true, err.Error(), "", false)
	}

	transaction.ID = string(uuid.GenerateUUID())

	err = transactiondb.CreateTransaction(&transaction, db)

	if err != nil {
		return result.New("", true, err.Error(), "", false)
	}

	return result.New(encoding.ConvertToJSON(transaction), false, "", "", true)
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
