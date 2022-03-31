package jwt

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	SecretKey = []byte("mysecret")
)

// Generate Token JWT by personal SecretKey
func GenerateToken(username string) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	signed, err := token.SignedString(SecretKey)
	if err != nil {
		log.Fatalf("Error for generate key for user:%s", username)
		return "", err
	} else {
		return signed, nil
	}
}

func ParseToken(tok string) (string, error) {
	token, err := jwt.Parse(tok, func(t *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		log.Fatalf("Can`t parse token error:%v", err)
		return "", err
	} else {
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			return claims["username"].(string), nil
		} else {
			log.Fatal("Can`t receive username from claims token")
			return "", err
		}
	}
}
