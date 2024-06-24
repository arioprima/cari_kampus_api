package pkg

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateToken(Payload interface{}, SecretJwtKey string, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(ttl).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"sub": Payload,
	})

	tokenString, err := token.SignedString([]byte(SecretJwtKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
