package participant

import (
	"crypto/sha256"

	"github.com/stoina/go-blockchain/blockchain/util/encoding"
)

// GetHashCode exported
// ...
func (participant *Participant) GetHashCode() []byte {
	return encoding.ConvertToSlice(sha256.Sum256([]byte(string(participant.ID) + " " + participant.EMail)))
}
