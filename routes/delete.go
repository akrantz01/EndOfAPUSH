package routes

import (
	"github.com/akrantz01/EndOfAPUSH/database"
	"github.com/akrantz01/EndOfAPUSH/util"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
)

func Delete(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Validate initial request on path and Authorization header
		path := mux.Vars(r)
		if path["username"] == "" {
			util.Responses.Error(w, http.StatusBadRequest, "must specify a username")
			return
		} else if r.Header.Get("Authorization") == "" {
			util.Responses.Error(w, http.StatusUnauthorized, "header 'Authorization' is required")
			return
		}

		// TODO: validate user authentication token

		// Check if user exists
		var user database.User
		db.Where("username = ?", path["username"]).First(&user)
		if user.ID == 0 {
			util.Responses.Error(w, http.StatusBadRequest, "user with specified username does not exist")
			return
		}

		// Delete the user
		db.Delete(&user)

		util.Responses.Success(w)
	}
}
