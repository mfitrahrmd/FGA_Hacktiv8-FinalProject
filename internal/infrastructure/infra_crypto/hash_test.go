package infra_crypto

import (
	"testing"
)

var hashedPassword string
var mypassword string = "mypassword"

func TestHashPassword(t *testing.T) {
	t.Logf("password to hash : %s", mypassword)

	hash, err := HashPassword(mypassword)
	if err != nil {
		t.Error(err.Error())
	}

	if hash == mypassword {
		t.Fail()
	}

	hashedPassword = hash

	t.Logf("hashed password : %s", hashedPassword)
}

func TestComparePassword(t *testing.T) {
	t.Logf("hashed password to compare : %s", hashedPassword)

	match := ComparePassword(mypassword, hashedPassword)

	if !match {
		t.Fail()
	}

	t.Log("password match")
}
