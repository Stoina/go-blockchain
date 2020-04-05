package uuid

import (
	"os/exec"
)

// GenerateUUID exported
// ...
func GenerateUUID() []byte {
	uuid, err := exec.Command("uuidgen").Output()

	if err != nil {
		return []byte{}
	}

	return uuid[:36]
}
