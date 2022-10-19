package infra_jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/config/env"
	"time"
)

type Claims struct {
	Payload any
	jwt.StandardClaims
}

func GenerateToken(key string, payload any) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		payload,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Millisecond * time.Duration(env.JWT_DURATION)).Unix(),
		},
	})

	signedToken, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ValidateToken(key string, jwtToken string) (any, error) {
	token, err := jwt.ParseWithClaims(jwtToken, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrInvalidKey
	}

	c, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("error parsing token type")
	}

	return c.Payload, nil
}
