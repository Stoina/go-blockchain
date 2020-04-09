package uuid

import (
	"encoding/hex"

	guuid "github.com/google/uuid"
)

// GenerateUUID exported
// ...
func GenerateUUID() string {
	uuid := guuid.New()
	return hex.EncodeToString(uuid[:])
}
