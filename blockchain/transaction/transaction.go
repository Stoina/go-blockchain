package transaction

import (
	participant "github.com/stoina/go-blockchain/blockchain/participant"
	uuid "github.com/stoina/go-blockchain/blockchain/uuid"
	value "github.com/stoina/go-blockchain/blockchain/value"
)

// Transaction exported
// ...
type Transaction struct {
	ID          string
	Transmitter *participant.Participant
	Receiver    *participant.Participant
	CreditValue *value.Value
}

// Create exported
// ...
func Create(transmitter *participant.Participant, receiver *participant.Participant, creditValue *value.Value) *Transaction {
	return &Transaction{
		ID:          string(uuid.GenerateUUID()),
		Transmitter: transmitter,
		Receiver:    receiver,
		CreditValue: creditValue}
}

// CreateWithID exported
// ...
func CreateWithID(id string, transmitter *participant.Participant, receiver *participant.Participant, creditValue *value.Value) *Transaction {
	return &Transaction{
		ID:          id,
		Transmitter: transmitter,
		Receiver:    receiver,
		CreditValue: creditValue}
}
