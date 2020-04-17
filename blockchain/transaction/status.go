package transaction

// Transaction status codes
//
const (
	StatusCreated   = 100
	StatusConfirmed = 200
	StatusValidated = 300
)

// Text to transaction status codes
//
var statusText = map[int]string{
	StatusCreated:   "Created",
	StatusConfirmed: "Confirmed",
	StatusValidated: "Validated"}

// StatusText exported
// ...
func StatusText(code int) string {
	return statusText[code]
}
