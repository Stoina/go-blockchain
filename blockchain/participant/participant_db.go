package participant

import (
	"encoding/json"
	"errors"
	"fmt"

	db "github.com/Stoina/go-database"
)

// ReadParticipantByID exported
// ...
func ReadParticipantByID(id string, dbConn *db.Connection) (*Participant, error) {

	result, err := dbConn.Query("select * from public.bc_participant where bcp_id = " + id)

	if err != nil {
		return nil, err
	}

	return createParticipantByResult(result)
}

// ReadParticipantByEMail exported
// ...
func ReadParticipantByEMail(email string, dbConn *db.Connection) (*Participant, error) {

	result, err := dbConn.Query("select * from public.bc_participant where bcp_email = " + email)

	if err != nil {
		return nil, err
	}

	return createParticipantByResult(result)
}

// ReadParticipants exported
// ...
func ReadParticipants(dbConn *db.Connection) ([]Participant, error) {
	result, err := dbConn.Query("select * from public.bc_participant")

	if err != nil {
		return nil, err
	}

	jsonDataString, err := result.ConvertDataToJSONString()

	if err != nil {
		return nil, err
	}

	participants := make([]Participant, result.RowCount)
	err = json.Unmarshal([]byte(jsonDataString), &participants)

	return participants, nil
}

// CreateOrUpdateParticipant exported
// ...
func (participant *Participant) CreateOrUpdateParticipant(dbConn *db.Connection) (string, error) {
	_, err := dbConn.CallProcedure("proc_create_or_update_participant", getProcedureParams(participant))
	return "", err
}

func createParticipantByResult(result *db.Result) (*Participant, error) {

	if result.RowCount == 1 {

		participant := &Participant{}

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

func getProcedureParams(participant *Participant) []interface{} {

	return []interface{}{
		participant.ID,
		participant.EMail,
		participant.Lastname,
		participant.Firstname,
		fmt.Sprintf("%x", participant.GetHashCode())}

}
