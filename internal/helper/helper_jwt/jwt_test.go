package helper_jwt

import (
	"encoding/json"
	"testing"
)

type Payload struct {
	Email string `json:"email"`
}

var mykey string = "mykey"
var mypayload = Payload{
	"rama1@gmail.com",
}
var createdtoken string

func TestCreateToken(t *testing.T) {
	token, err := GenerateToken(mykey, mypayload)
	if err != nil {
		t.Fail()
		t.Log(err.Error())
	}

	createdtoken = token
	t.Log(token)
}

func TestValidateToken(t *testing.T) {
	payload, err := ValidateToken(mykey, createdtoken)
	if err != nil {
		t.Fail()
		t.Log(err.Error())
	}

	p := Payload{}
	m, _ := json.Marshal(payload)
	json.Unmarshal(m, &p)

	t.Logf("%+v", p)
}
