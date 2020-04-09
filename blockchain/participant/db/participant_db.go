package participantdb

import (
	"encoding/json"
	"errors"
	"fmt"

	participant "github.com/stoina/go-blockchain/blockchain/participant"
	db "github.com/stoina/go-database"
)

// ReadParticipantByID exported
// ...
func ReadParticipantByID(id string, dbConn *db.Connection) (*participant.Participant, error) {

	result, err := dbConn.Query("select * from public.bc_participant where bcp_id = " + id)

	if err != nil {
		return nil, err
	}

	return newParticipantByResult(result)
}

// ReadParticipantByEMail exported
// ...
func ReadParticipantByEMail(email string, dbConn *db.Connection) (*participant.Participant, error) {

	result, err := dbConn.Query("select * from public.bc_participant where bcp_email = " + email)

	if err != nil {
		return nil, err
	}

	return newParticipantByResult(result)
}

// ReadParticipants exported
// ...
func ReadParticipants(dbConn *db.Connection) ([]participant.Participant, error) {
	result, err := dbConn.Query("select * from public.bc_participant")

	if err != nil {
		return nil, err
	}

	jsonDataString, err := result.ConvertDataToJSONString()

	if err != nil {
		return nil, err
	}

	participants := make([]participant.Participant, result.RowCount)
	err = json.Unmarshal([]byte(jsonDataString), &participants)

	return participants, nil
}

// CreateOrUpdateParticipant exported
// ...
func CreateOrUpdateParticipant(participant *participant.Participant, dbConn *db.Connection) error {
	_, err := dbConn.CallProcedure("proc_create_or_update_participant", getProcedureParams(participant))
	return err
}

func newParticipantByResult(result *db.Result) (*participant.Participant, error) {

	if result.RowCount == 1 {

		participant := &participant.Participant{}

		jsonDataString, err := result.ConvertDataToJSONString()

		if err != nil {
			return nil, err
		}

		json.Unmarshal([]byte(jsonDataString), &participant)

		return participant, nil

	} else if result.RowCount > 1 {
		return nil, errors.New("Multiple rows found for E-Mail address")
	}

	return nil, errors.New("Now row found for E-Mail address")

}

func getProcedureParams(participant *participant.Participant) []interface{} {

	return []interface{}{
		participant.ID,
		participant.EMail,
		participant.Lastname,
		participant.Firstname,
		fmt.Sprintf("%x", participant.GetHashCode())}

}
