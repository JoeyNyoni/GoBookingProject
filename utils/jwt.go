package utils

import (
	"errors"
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

func VerifyJWTToken(tokenString string) (int64, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method to prevent security issues
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secretKey), nil // Use the defined secret key
	})

	if err != nil {
		return 0, errors.New("Could not verify JWT token")
	}

	isTokenValid := parsedToken.Valid
	if !isTokenValid {
		return 0, errors.New("Invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("Invalid claims from token")
	}

	//email := claims["email"].(string)
	userId := int64(claims["userId"].(float64)) // JWT claims are typically float64 for numeric values
	return userId, nil
}
