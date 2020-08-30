package participant

import (
	uuid "github.com/Stoina/go-blockchain/blockchain/uuid"
)

// Participant exported
// ...
type Participant struct {
	ID        string `json:"bcp_id"`
	EMail     string `json:"bcp_email"`
	Lastname  string `json:"bcp_lastname"`
	Firstname string `json:"bcp_firstname"`
}

// New exported
// ...
func New(email string, lastname string, firstname string) *Participant {
	return &Participant{
		ID:        string(uuid.GenerateUUID()),
		EMail:     email,
		Lastname:  lastname,
		Firstname: firstname}
}

// NewWithID exported
// ...
func NewWithID(id string, email string, lastname string, firstname string) *Participant {
	return &Participant{
		ID:        id,
		EMail:     email,
		Lastname:  lastname,
		Firstname: firstname}
}
