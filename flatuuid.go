package flatuuid

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

// Compact makes given UUID more compact
func Compact(srcUUID string) (string, error) {
	if len(strings.Trim(srcUUID, " ")) == 0 {
		return "", errors.New("Empty UUID string")
	}

	parsedUUID, err := uuid.Parse(srcUUID)
	if err != nil {
		return "", err
	}

	compact := strings.ReplaceAll(parsedUUID.String(), "-", "")
	return compact, nil
}
