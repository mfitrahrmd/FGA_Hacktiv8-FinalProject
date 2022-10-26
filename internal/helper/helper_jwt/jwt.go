package helper_jwt

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/config/env"
)

type TokenPayload struct {
	Id       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type Claims struct {
	Payload any
	jwt.StandardClaims
}

func GenerateToken(key string, payload any) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		payload,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(env.JWT_DURATION)).Unix(),
		},
	})

	signedToken, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ValidateToken(key string, jwtToken string) (*TokenPayload, error) {
	token, err := jwt.ParseWithClaims(jwtToken, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(fmt.Sprintf("error token signing method : %v", t.Header["alg"]))
		}

		return []byte(key), nil
	})
	if err != nil {
		vErr := err.(*jwt.ValidationError)
		fmt.Println(vErr.Inner)
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrInvalidKey
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("error parsing token type")
	}

	var tokenPayload TokenPayload

	marshal, _ := json.Marshal(claims.Payload)

	json.Unmarshal(marshal, &tokenPayload)

	return &tokenPayload, nil
}
