package handlers

import (
	"encoding/json"
	"net/http"

	"BIKE-SHARING-SERVICE/internal/config"
	"BIKE-SHARING-SERVICE/internal/db/models"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(cfg *config.Config, sessionStore sessions.Store, db *db.Database) http.HandlerFunc {
	// swagger:operation POST /login
	//
	// Include Documentation
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: body
	//   in: body
	//   schema:
	//   	"$ref": "#/definitions/login"
	// responses:
	//   '200':
	//     description: user response

	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type response struct {
		User *models.User `json:"user"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()

		var req request
		err := decoder.Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := db.GetUserByEmail(req.Email)
		if err != nil {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
		if err != nil {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		// Create a new session and save the user ID in it
		session, err := sessionStore.New(r, cfg.SessionName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		session.Values["userID"] = user.ID
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := response{User: user}
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResp)
	}
}
