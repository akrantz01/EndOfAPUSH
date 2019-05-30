package messages

import (
	"github.com/akrantz01/EndOfAPUSH/database"
	"github.com/akrantz01/EndOfAPUSH/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"net/http"
)

func List(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Validate initial request on headers
		if r.Header.Get("Authorization") == "" {
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
		db.Set("gorm:auto_preload", true).Where("id = ?", claims["sub"]).First(&user)

		// Assemble message into map for JSON
		var in []map[string]interface{}
		var out []map[string]interface{}

		// Handle inbox messages
		for _, msg := range user.Inbox {
			j := make(map[string]interface{})

			// Get name of sender or [deleted]
			var from database.User
			db.Where("id = ?", msg.FromID).First(&from)
			if from.ID == 0 {
				from.Name = "[deleted]"
			}

			j["id"] = msg.ID
			j["subject"] = msg.Subject
			j["from"] = from.Name

			in = append(in, j)
		}

		// Handle outbox messages
		for _, msg := range user.Outbox {
			j := make(map[string]interface{})

			// Get name of recipient or [deleted]
			var to database.User
			db.Where("id = ?", msg.FromID).First(&to)
			if to.ID == 0 {
				to.Name = "[deleted]"
			}

			j["id"] = msg.ID
			j["subject"] = msg.Subject
			j["from"] = to.Name

			out = append(out, j)
		}

		// Respond with different data based on query parameters
		switch r.URL.Query().Get("type") {
		case "in":
			if in == nil {
				in = []map[string]interface{}{}
			}
			util.Responses.SuccessWithData(w, in)
		case "out":
			if out == nil {
				out = []map[string]interface{}{}
			}
			util.Responses.SuccessWithData(w, out)
		case "":
			util.Responses.SuccessWithData(w, append(in, out...))
		default:
			util.Responses.Error(w, http.StatusBadRequest, "query parameter 'type' must be one of 'in' or 'out' or undefined")
			return
		}
	}
}
