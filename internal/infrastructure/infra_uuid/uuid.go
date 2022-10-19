package infra_uuid

import (
	"github.com/google/uuid"
	"strings"
)

func GenerateUUID(prefix ...string) string {
	return prefix[0] + "-" + strings.Replace(uuid.New().String(), "-", "", -1)
}
