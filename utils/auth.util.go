package utils

import (
	"os"
	"strconv"

	"example.com/config"
	"example.com/models"
	"golang.org/x/crypto/bcrypt"
  jwt "github.com/golang-jwt/jwt/v5"
)

func verifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func generateToken(userID uint) (token string, err error) {
  tokenLifespan, err := strconv.Atoi(os.Getenv("TOKEN_LIFESPAN"))

  if err != nil {
    return "", err
  }

}

func LoginCheck(username, password string) (string, error) {
	DB := config.ConnectDB()
	var u models.User

	// Find user by username
	DB.First(&u, "username = ?", username)

  // Verify password
  if err := verifyPassword(password, u.Password); err != nil {
    return "", err
  }

  token, err := generateToken(u.ID)

  if err != nil {
    return "", err
  }

  return token, nil
}
