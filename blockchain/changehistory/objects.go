package changehistory

// Object Ids
//
const (
	ParticipantTable = 1
	BlockTable       = 2
	TransactionTable = 3
)

// Name to objects
//
var objectNames = map[int]string{
	ParticipantTable: "Participant",
	BlockTable:       "Block",
	TransactionTable: "Transaction"}

// ObjectName exported
// ...
func ObjectName(id int) string {
	return objectNames[id]
}
