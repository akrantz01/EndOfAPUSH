package user

import (
	"fmt"
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

		// TODO: validate user authentication token

		// Find all users like given username
		var users []database.User
		db.Where("username LIKE ?", r.URL.Query().Get("username") + "%").Find(&users)
		fmt.Println(users)

		// Create username -> name map
		response := make(map[string]string)
		for _, user := range users {
			response[user.Username] = user.Name
		}

		util.Responses.SuccessWithData(w, response)
	}
}
