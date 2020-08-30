package transaction

import (
	"crypto/sha256"

	"github.com/Stoina/go-blockchain/blockchain/util/encoding"
)

// GetHashCode exported
// ...
func (transaction *Transaction) GetHashCode() []byte {

	transmitterHashCode := transaction.Transmitter.GetHashCode()
	receiverHashCode := transaction.Receiver.GetHashCode()
	creditValueHashCode := transaction.CreditValue.GetHashCode()

	hashCodeBytes := []byte{}

	hashCodeBytes = append(transmitterHashCode, receiverHashCode...)
	hashCodeBytes = append(hashCodeBytes, creditValueHashCode...)

	return encoding.ConvertToSlice(sha256.Sum256(hashCodeBytes))
}
