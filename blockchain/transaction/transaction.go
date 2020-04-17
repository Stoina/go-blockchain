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
	Status      int
	CreditValue *value.Value
}

// New exported
// ...
func New(transmitter *participant.Participant, receiver *participant.Participant, creditValue *value.Value) *Transaction {
	return &Transaction{
		ID:          string(uuid.GenerateUUID()),
		Transmitter: transmitter,
		Receiver:    receiver,
		Status:      StatusCreated,
		CreditValue: creditValue}
}

// NewWithID exported
// ...
func NewWithID(id string, transmitter *participant.Participant, receiver *participant.Participant, creditValue *value.Value) *Transaction {
	return &Transaction{
		ID:          id,
		Transmitter: transmitter,
		Receiver:    receiver,
		Status:      StatusCreated,
		CreditValue: creditValue}
}
