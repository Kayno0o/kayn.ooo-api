package auth

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Make sure that the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Return the secret key used for signing the token
		return []byte("secret_key"), nil // replace with your own secret key
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
