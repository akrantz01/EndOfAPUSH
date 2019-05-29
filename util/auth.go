package util

import (
	"encoding/base64"
	"fmt"
	"github.com/akrantz01/EndOfAPUSH/database"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

func ValidateJWT(tokenString string, db *gorm.DB) (*jwt.Token, error) {
	// Retrieve token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, e error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		} else if _, ok := token.Header["kid"]; !ok {
			return nil, fmt.Errorf("unable to find key id in token")
		}

		// Get signing key from database
		var t database.Token
		db.Where("id = ?", token.Header["kid"]).First(&t)
		if t.SigningKey == "" {
			return nil, fmt.Errorf("unable to find signing key for token: %v", token.Header["kid"])
		}

		// Decode signing key from database
		signingKey, err := base64.StdEncoding.DecodeString(t.SigningKey)
		if err != nil {
			return nil, fmt.Errorf("unable to decode signing key: %v", err)
		}

		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}

	// Check if token is valid
	if !token.Valid {
		return nil, fmt.Errorf("token is invalid: %v", err)
	}

	return token, nil
}
