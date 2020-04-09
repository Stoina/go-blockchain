package value

import (
	"crypto/sha256"

	"github.com/stoina/go-blockchain/blockchain/encoding"
	binaryencoding "github.com/stoina/go-blockchain/blockchain/encoding"
)

// GetHashCode exported
// ...
func (value *Value) GetHashCode() []byte {
	return encoding.ConvertToSlice(sha256.Sum256(binaryencoding.Float64ToBytes(value.Value)))
}
