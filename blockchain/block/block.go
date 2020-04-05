package block

import (
	"github.com/Stoina/go-blockchain/blockchain/transaction"
)

// Block exported
// ...
type Block struct {
	ID            string
	Transactions  []transaction.Transaction
	PreviousBlock *Block
}
