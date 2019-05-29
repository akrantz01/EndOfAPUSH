package messages

import (
	"encoding/json"
	"github.com/akrantz01/EndOfAPUSH/database"
	"github.com/akrantz01/EndOfAPUSH/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"net/http"
)

func Create(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Validate initial request on headers
		if r.Header.Get("Content-Type") != "application/json" {
			util.Responses.Error(w, http.StatusBadRequest, "header 'Content-Type' must be 'application/json'")
			return
		} else if r.Header.Get("Authorization") == "" {
			util.Responses.Error(w, http.StatusUnauthorized, "header 'Authorization' is required")
			return
		}

		// Validate JSON body
		var body struct{
			To        string
			Algorithm int
			Subject   string
			Message   string
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			util.Responses.Error(w, http.StatusBadRequest, "unable to decode JSON: " + err.Error())
			return
		} else if body.To == "" || body.Algorithm == 0 || body.Subject == "" || body.Message == "" {
			util.Responses.Error(w, http.StatusBadRequest, "fields 'to', 'algorithm', 'subject', and 'password' are required")
			return
		} else if body.Algorithm > 8 || body.Algorithm < 1 {
			// SIGABA: 1
			// DES: 2
			// TripleDES: 3
			// AES: 4
			// Diffie-Hellman DES: 5
			// Diffie-Hellman TripleDES: 6
			// Diffie-Hellman AES: 7
			// RSA: 8
			util.Responses.Error(w, http.StatusBadRequest, "field 'algorithm' must be between 1 and 8")
			return
		}

		// Verify JWT from header
		token, err := util.ValidateJWT(r.Header.Get("Authorization"), db)
		if err != nil {
			util.Responses.Error(w, http.StatusUnauthorized, "invalid token: " + err.Error())
			return
		}

		// Get claims from JWT
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			util.Responses.Error(w, http.StatusInternalServerError, "invalid claims format")
			return
		}

		// Get user from JWT in header
		var user database.User
		db.Where("id = ?", claims["sub"]).First(&user)

		// Check that to user exists
		var toUser database.User
		db.Where("username = ?", body.To).First(&toUser)
		if toUser.ID == 0 {
			util.Responses.Error(w, http.StatusBadRequest, "user specified in 'to' field does not exist")
			return
		}

		// Check that not sending to self
		if toUser.ID == user.ID {
			util.Responses.Error(w, http.StatusBadRequest, "cannot send message to self")
			return
		}

		// Create message for database
		msg := &database.Message{
			ToID: toUser.ID,
			FromID: user.ID,
			Subject: body.Subject,
			Message: body.Message,
			Algorithm: uint(body.Algorithm),
		}
		db.NewRecord(msg)
		db.Create(&msg)

		util.Responses.Success(w)
	}
}
