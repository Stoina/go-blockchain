package changehistory

// Transaction status codes
//
const (
	StatusCreated = 1
	StatusUpdated = 2
	StatusDeleted = 3
)

// Text to transaction status codes
//
var statusTexts = map[int]string{
	StatusCreated: "Created",
	StatusUpdated: "Updated",
	StatusDeleted: "Deleted"}

// StatusText exported
// ...
func StatusText(code int) string {
	return statusTexts[code]
}
