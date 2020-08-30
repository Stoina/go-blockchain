package transactiondb

import (
	"fmt"

	"github.com/Stoina/go-blockchain/blockchain/transaction"
	db "github.com/Stoina/go-database"
)

// ReadTransactionByID exported
// ...
func ReadTransactionByID(id string, dbConn *db.Connection) (*transaction.Transaction, error) {
	return nil, nil
}

// CreateTransaction exported
// ...
func CreateTransaction(transaction *transaction.Transaction, dbConn *db.Connection) error {
	_, err := dbConn.Insert(getInsertStatement(transaction))
	return err
}

func getInsertStatement(transaction *transaction.Transaction) *db.InsertStatement {
	return db.NewInsertStatement("bc_transaction", []string{"bct_id", "bct_transmitter_id", "bct_receiver_id", "bct_creditvalue", "bct_status", "bct_hash"}, []interface{}{
		transaction.ID,
		transaction.Transmitter.ID,
		transaction.Receiver.ID,
		transaction.CreditValue.Value,
		transaction.Status,
		fmt.Sprintf("%x", transaction.GetHashCode())})
}
