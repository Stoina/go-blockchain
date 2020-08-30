package value

import (
	"crypto/sha256"

	"github.com/Stoina/go-blockchain/blockchain/util/encoding"
)

// GetHashCode exported
// ...
func (value *Value) GetHashCode() []byte {
	return encoding.ConvertToSlice(sha256.Sum256(encoding.Float64ToBytes(value.Value)))
}
