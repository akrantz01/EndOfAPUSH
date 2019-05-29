package messages

import (
	"github.com/akrantz01/EndOfAPUSH/database"
	"github.com/akrantz01/EndOfAPUSH/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
)

func Read(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Validate initial request on headers and path
		path := mux.Vars(r)
		if path["id"] == "" {
			util.Responses.Error(w, http.StatusBadRequest, "a message id must be specified in path")
			return
		} else if r.Header.Get("Authorization") == "" {
			util.Responses.Error(w, http.StatusUnauthorized, "header 'Authorization' is required")
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

		// Check that message exists
		var message database.Message
		db.Where("id = ?", path["id"]).First(&message)
		if message.ID == 0 {
			util.Responses.Error(w, http.StatusBadRequest, "specified message does not exist")
			return
		}

		// Get to username and from username

		var (
			otherUser database.User
			toName string
			fromName string
		)
		switch user.ID {
		case message.ToID:
			db.Where("id = ?", message.FromID).First(&otherUser)
			toName = user.Name
			fromName = otherUser.Name
		case message.FromID:
			db.Where("id = ?", message.ToID).First(&otherUser)
			fromName = user.Name
			toName = otherUser.Name
		}
		if otherUser.ID == 0 {
			util.Responses.Error(w, http.StatusInternalServerError, "from user no longer exists")
			return
		}

		// Assemble message into map for JSON
		response := make(map[string]interface{})
		response["subject"] = message.Subject
		response["message"] = message.Message
		response["algorithm"] = message.Algorithm
		response["to"] = toName
		response["from"] = fromName

		util.Responses.SuccessWithData(w, response)
	}
}
