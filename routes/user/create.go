package user

import (
	"encoding/json"
	"github.com/akrantz01/EndOfAPUSH/database"
	"github.com/akrantz01/EndOfAPUSH/util"
	"github.com/jinzhu/gorm"
	"gopkg.in/hlandau/passlib.v1"
	"net/http"
)

func Create(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
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
			Name     string
			Username string
			Password string
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			util.Responses.Error(w, http.StatusBadRequest, "unable to decode JSON: " + err.Error())
			return
		} else if body.Name == "" || body.Username == "" || body.Password == "" {
			util.Responses.Error(w, http.StatusBadRequest, "fields 'name', 'username', and 'password' are required")
			return
		}

		// Check if username is already taken
		var user database.User
		db.Where("username = ?", body.Username).First(&user)
		if user.ID != 0 {
			util.Responses.Error(w, http.StatusBadRequest, "username is already in use")
			return
		}

		// Hash the password for security
		hash, err := passlib.Hash(body.Password)
		if err != nil {
			util.Responses.Error(w, http.StatusInternalServerError, "failed to hash password: " + err.Error())
			return
		}

		// Assemble user from body
		user.Name = body.Name
		user.Username = body.Username
		user.Password = hash

		// Save to database
		db.NewRecord(user)
		db.Save(&user)

		util.Responses.Success(w)
	}
}
