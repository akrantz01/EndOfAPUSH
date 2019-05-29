package auth

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/akrantz01/EndOfAPUSH/database"
	"github.com/akrantz01/EndOfAPUSH/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"gopkg.in/hlandau/passlib.v1"
	"net/http"
	"time"
)

func Login(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Validate initial request on Content-Type header and body
		if r.Header.Get("Content-Type") != "application/json" {
			util.Responses.Error(w, http.StatusBadRequest, "content must be JSON")
			return
		} else if r.Body == nil {
			util.Responses.Error(w, http.StatusBadRequest, "body must exist")
			return
		}

		// Validate JSON body
		var body struct{
			Username string
			Password string
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			util.Responses.Error(w, http.StatusBadRequest, "unable to decode JSON: " + err.Error())
			return
		} else if body.Username == "" || body.Password == "" {
			util.Responses.Error(w, http.StatusBadRequest, "fields 'username' and 'password' are required")
			return
		}

		// Check if user exists
		var user database.User
		db.Where("username = ?", body.Username).First(&user)
		if user.ID == 0 {
			util.Responses.Error(w, http.StatusUnauthorized, "invalid username or password")
			return
		}

		// Validate password
		newHash, err := passlib.Verify(body.Password, user.Password)
		if err != nil {
			util.Responses.Error(w, http.StatusUnauthorized, "invalid username or password")
			return
		}

		// Update password hash if needed
		if  newHash != "" {
			user.Password = newHash
			db.Save(&user)
		}

		// Create signing key and write to database
		signingKey := make([]byte, 128)
		if _, err := rand.Read(signingKey); err != nil {
			util.Responses.Error(w, http.StatusInternalServerError, "failed to generate JWT signing key: " + err.Error())
			return
		}
		storedToken := &database.Token{
			UserId: user.ID,
			SigningKey: base64.StdEncoding.EncodeToString(signingKey),
		}
		db.NewRecord(storedToken)
		db.Create(&storedToken)

		// Create claims for JWT
		claims := &jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + (60*60*24),
			Issuer: "endofapush",
			IssuedAt: time.Now().Unix(),
			Subject: fmt.Sprint(user.ID),
		}

		// Generate token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		token.Header["kid"] = storedToken.ID

		// Sign token
		signed, err := token.SignedString(signingKey)
		if err != nil {
			util.Responses.Error(w, http.StatusInternalServerError, "failed to sign JWT: " + err.Error())
			return
		}

		util.Responses.SuccessWithData(w, map[string]string{"token": signed})
	}
}
