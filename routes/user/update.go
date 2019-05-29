package user

import (
	"encoding/json"
	"github.com/akrantz01/EndOfAPUSH/database"
	"github.com/akrantz01/EndOfAPUSH/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"gopkg.in/hlandau/passlib.v1"
	"net/http"
)

func Update(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Validate initial request on headers and path
		path := mux.Vars(r)
		if path["username"] == "" {
			util.Responses.Error(w, http.StatusBadRequest, "a username must be specified")
			return
		} else if r.Header.Get("Content-Type") != "application/json" {
			util.Responses.Error(w, http.StatusBadRequest, "header 'Content-Type' must be JSON")
			return
		} else if r.Header.Get("Authorization") == "" {
			util.Responses.Error(w, http.StatusUnauthorized, "header 'Authorization' is required")
			return
		}

		// Validate JSON body
		var body struct {
			Name     string
			Username string
			Password string
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			util.Responses.Error(w, http.StatusBadRequest, "unable to decode JSON: "+err.Error())
			return
		} else if body.Name == "" && body.Username == "" && body.Password == "" {
			// Nothing changed, return success
			util.Responses.Success(w)
			return
		}

		// Verify JWT from headers
		token, err := util.ValidateJWT(r.Header.Get("Authorization"), db)
		if err != nil {
			util.Responses.Error(w, http.StatusUnauthorized, "invalid token: "+err.Error())
			return
		}

		// Check if user exists
		var user database.User
		db.Where("username = ?", path["username"]).First(&user)
		if user.ID == 0 {
			util.Responses.Error(w, http.StatusBadRequest, "user with specified username does not exist")
			return
		}

		// Retrieve token claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			util.Responses.Error(w, http.StatusInternalServerError, "invalid claims format")
			return
		}

		// Retrieve user specified in token
		var tokenUser database.User
		db.Where("id = ?", claims["sub"]).First(&tokenUser)
		if tokenUser.ID == 0 {
			util.Responses.Error(w, http.StatusUnauthorized, "no user exists at id: "+claims["sub"].(string))
			return

		// Check that specified user and token match
		} else if tokenUser.ID != user.ID {
			util.Responses.Error(w, http.StatusUnauthorized, "specified user and token do not match")
			return
		}

		// Update values if in body
		if body.Name != "" {
			user.Name = body.Name
		}
		if body.Username != "" {
			user.Username = body.Username
		}
		if body.Password != "" {
			hash, err := passlib.Hash(body.Password)
			if err != nil {
				util.Responses.Error(w, http.StatusInternalServerError, "failed to hash password: " + err.Error())
				return
			}
			user.Password = hash
		}

		// Save update
		db.Save(&user)

		util.Responses.Success(w)
	}
}
