package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("yourSecret")

func GenrateToken(email string, userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString(secretKey)

}

func ValidateToken(tokenString string) (int64, error) {
	JWTtoken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unvalid token")
		}
		return secretKey, nil
	})

	if err != nil {
		return 0, err
	}

	validToken := JWTtoken.Valid

	if !validToken {
		return 0, errors.New("token is not valid")
	}

	comma, ok := JWTtoken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, err
	}

	// email := comma["email"].(string)
	userID := int64(comma["userID"].(float64))
	return userID, nil
}
