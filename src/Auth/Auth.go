package auth

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	entity "kayn.ooo/api/src/Entity"
	repository "kayn.ooo/api/src/Repository"
)

type JWT struct {
	Token string `json:"token"`
}

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(user *entity.User) (*JWT, error) {
	// Load the secret key from the .env file
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	secretKey := []byte(os.Getenv("SECRET_KEY"))

	// Create a new JWT token with the user ID and expiration time
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Sign the token with the secret key and return a JWT struct with the token string
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return nil, err
	}
	return &JWT{Token: tokenString}, nil
}

func Authenticate(email, password string) (*entity.User, error) {
	var user entity.User

	result := repository.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, err
	}

	return &user, nil
}

func GetUserFromToken(tokenString string) (*entity.User, error) {
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	// Load the secret key from the .env file
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	secretKey := []byte(os.Getenv("SECRET_KEY"))

	// Parse the JWT token string
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		userID := claims.UserID

		var user entity.User
		err := repository.FindByID(userID, &user)
		if err != nil {
			return nil, err
		}

		return &user, nil
	} else {
		return nil, errors.New("invalid token")
	}
}
