package utils

import (
	"os"
	"strconv"
	"time"

	"example.com/config"
	"example.com/models"
	jwt "github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func verifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func generateToken(userID uint) (string, error) {
	tokenLifespan, err := strconv.Atoi(os.Getenv("TOKEN_LIFESPAN"))

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized": true,
		"userID":     userID,
		"exp":        time.Now().Add(time.Hour * time.Duration(tokenLifespan)).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("SECRET_JWT")))
}

func LoginCheck(username, password string) (string, error) {
	DB := config.ConnectDB()
	var u models.User

	// Find user by username
	DB.First(&u, "username = ?", username)

	// Verify password
	if err := verifyPassword(password, u.Password); err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := generateToken(u.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}
