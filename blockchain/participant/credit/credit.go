package credit

import (
	"github.com/Stoina/go-blockchain/blockchain/participant"
	"github.com/Stoina/go-blockchain/blockchain/value"
)

// Credit exported
// ...
type Credit struct {
	Participant *participant.Participant
	Value       *value.Value
}

// New exported
// ...
func New(participant *participant.Participant, value *value.Value) *Credit {
	return &Credit{
		Participant: participant,
		Value:       value}
}
