package pkg

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
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

func ValidateToken(token string, signedJWTKey string) (interface{}, error) {
	tok, _ := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", jwtToken.Header["alg"])
		}
		return []byte(signedJWTKey), nil
	})

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		logrus.Error("token is invalid")
	}
	return claims["sub"], nil
}
