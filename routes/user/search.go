package user

import (
	"github.com/akrantz01/EndOfAPUSH/database"
	"github.com/akrantz01/EndOfAPUSH/util"
	"github.com/jinzhu/gorm"
	"net/http"
)

func Search(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Validate initial request on query parameters
		if len(r.URL.RawQuery) == 0 {
			util.Responses.Error(w, http.StatusBadRequest, "query parameters must exist")
			return
		} else if r.URL.Query().Get("username") == "" {
			util.Responses.Error(w, http.StatusBadRequest, "query parameter 'username' is required")
			return
		} else if r.Header.Get("Authorization") == "" {
			util.Responses.Error(w, http.StatusUnauthorized, "header 'Authorization' is required")
			return
		}

		// Verify JWT from headers
		_, err := util.ValidateJWT(r.Header.Get("Authorization"), db)
		if err != nil {
			util.Responses.Error(w, http.StatusUnauthorized, "invalid token: " + err.Error())
			return
		}

		// Find all users like given username
		var users []database.User
		db.Where("username LIKE ?", r.URL.Query().Get("username") + "%").Find(&users)

		// Create array of usernames and names
		var response []map[string]string
		for _, user := range users {
			response = append(response, map[string]string{"name": user.Name, "username": user.Username})
		}

		util.Responses.SuccessWithData(w, response)
	}
}
