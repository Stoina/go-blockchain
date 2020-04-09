package repository

import (
	"encoding/json"

	participant "github.com/stoina/go-blockchain/blockchain/participant"
	participantdb "github.com/stoina/go-blockchain/blockchain/participant/db"
	"github.com/stoina/go-blockchain/blockchain/util/encoding"
	"github.com/stoina/go-blockchain/blockchain/uuid"
	"github.com/stoina/go-blockchain/cmd/server/bcserver/repository/result"
	db "github.com/stoina/go-database"
)

// ParticipantRepository export
// ...
type ParticipantRepository struct {
}

// CreateParticipantRepository exported
// ...
func CreateParticipantRepository() *ParticipantRepository {
	return &ParticipantRepository{}
}

// Create exported
// ...
func (pr *ParticipantRepository) Create(jsonString string, db *db.Connection) *result.RepositoryResult {

	var participant participant.Participant
	err := json.Unmarshal([]byte(jsonString), &participant)

	if err != nil {
		return result.Create("", true, err.Error(), "", false)
	}

	participant.ID = string(uuid.GenerateUUID())

	err = participantdb.CreateOrUpdateParticipant(&participant, db)

	if err != nil {
		return result.Create("", true, err.Error(), "", false)
	}

	return result.Create(encoding.ConvertToJSON(participant), false, "", "", true)
}

// Update exported
// ...
func (pr *ParticipantRepository) Update(jsonString string, db *db.Connection) *result.RepositoryResult {
	return nil
}

// Delete exported
// ...
func (pr *ParticipantRepository) Delete(jsonString string, db *db.Connection) *result.RepositoryResult {
	return nil
}

// ReadAll exported
// ...
func (pr *ParticipantRepository) ReadAll(db *db.Connection) *result.RepositoryResult {
	participants, err := participantdb.ReadParticipants(db)

	if err != nil {
		return result.Create("", true, err.Error(), "", false)
	}

	return result.Create(encoding.ConvertToJSON(participants), false, "", "", true)
}

// ReadByID exported
// ...
func (pr *ParticipantRepository) ReadByID(id string, db *db.Connection) *result.RepositoryResult {
	return nil
}

// ReadByProperty exported
// ...
func (pr *ParticipantRepository) ReadByProperty(propertyName string, value interface{}, db *db.Connection) *result.RepositoryResult {
	return nil
}
