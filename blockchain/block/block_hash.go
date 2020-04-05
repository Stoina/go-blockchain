package block

import (
	"crypto/sha256"

	"github.com/Stoina/go-blockchain/blockchain/encoding"
)

// GetHashCode exported
// ...
func (block *Block) GetHashCode() []byte {

	hashCodeBytes := []byte{}

	for _, transaction := range block.Transactions {
		hashCodeBytes = append(hashCodeBytes, transaction.GetHashCode()...)
	}

	hashCodeBytes = append(hashCodeBytes, block.PreviousBlock.GetHashCode()...)

	return encoding.ConvertToSlice(sha256.Sum256(hashCodeBytes))
}
