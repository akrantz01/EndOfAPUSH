package user

import (
	"github.com/akrantz01/EndOfAPUSH/database"
	"github.com/akrantz01/EndOfAPUSH/util"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
)

func Read(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Validate initial request on path and Authorization header
		path := mux.Vars(r)
		if path["username"] == "" {
			util.Responses.Error(w, http.StatusBadRequest, "a username must be specified")
			return
		} else if r.Header.Get("Authorization") == "" {
			util.Responses.Error(w, http.StatusUnauthorized, "header 'Authorization' is required")
			return
		}

		// Verify JWT from headers
		token, err := util.ValidateJWT(r.Header.Get("Authorization"), db)
		if err != nil {
			util.Responses.Error(w, http.StatusUnauthorized, "invalid token: " + err.Error())
			return
		}

		// Check if user exists
		var user database.User
		db.Where("username = ?", path["username"]).First(&user)
		if user.ID == 0 {
			util.Responses.Error(w, http.StatusBadRequest, "user with specified username does not exist")
			return
		}

		// Check that specified user and token match
		if status, err := util.CheckJWTUsers(token, user, db); err != nil {
			util.Responses.Error(w, http.StatusBadRequest, err.Error())
			return
		} else if !status {
			util.Responses.Error(w, http.StatusUnauthorized, "specified user and token do not match")
		}

		// Assemble user into map for JSON
		response := make(map[string]string)
		response["name"] = user.Name
		response["username"] = user.Username

		util.Responses.SuccessWithData(w, response)
	}
}
