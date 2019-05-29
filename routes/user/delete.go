package user

import (
	"github.com/akrantz01/EndOfAPUSH/database"
	"github.com/akrantz01/EndOfAPUSH/util"
	"github.com/dgrijalva/jwt-go"
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
			util.Responses.Error(w, http.StatusUnauthorized, "no user exists at id: " + claims["sub"].(string))
			return
		}

		// Check that specified user and token match
		if tokenUser.ID != user.ID {
			util.Responses.Error(w, http.StatusUnauthorized, "specified user and token do not match")
			return
		}

		// Delete the user
		db.Delete(&user)

		util.Responses.Success(w)
	}
}
