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

// Create exported
// ...
func Create(participant *participant.Participant, value *value.Value) *Credit {
	return &Credit{
		Participant: participant,
		Value:       value}
}
