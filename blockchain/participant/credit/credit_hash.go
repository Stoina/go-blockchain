package credit

import (
	"crypto/sha256"

	"github.com/stoina/go-blockchain/blockchain/util/encoding"
)

// GetHashCode exported
// ...
func (credit *Credit) GetHashCode() []byte {

	participantHashCode := credit.Participant.GetHashCode()
	valueHashCode := credit.Value.GetHashCode()

	hashCodeBytes := []byte{}
	hashCodeBytes = append(participantHashCode, valueHashCode...)

	return encoding.ConvertToSlice(sha256.Sum256(hashCodeBytes))
}
