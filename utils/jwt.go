package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "your_secret_key" // dummy secret key, replace with your actual secret key

func GenerateJWTToken(email string, userId int64) (string, error) {
	// Define token claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(), // Token expires in 2 hours
	})

	return token.SignedString([]byte(secretKey)) // Use the defined secret key
}
