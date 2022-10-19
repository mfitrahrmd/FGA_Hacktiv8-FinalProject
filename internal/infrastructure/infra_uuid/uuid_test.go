package infra_uuid

import "testing"

func TestGenerateUUID(t *testing.T) {
	userid := GenerateUUID("user")

	t.Log(userid)
}
