package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

type RentRequest struct {
	BicycleID uint `json:"bicycle_id"`
}

// // swagger:operation POST bicycles/rent/{bicycleID} rentBicycle
// //
// // ---
// // summary: Rent a bicycle.
// // description: Rent a bicycle with the given ID.
// // parameters:
// //   - name: bicycleID
// //     in: path
// //     description: ID of the bicycle to rent.
// //     required: true
// //     schema:
// //     type: integer
// //     format: int64
// //
// // responses:
// //
// //	"200":
// //	  description: Success
// //	  content:
// //	    application/json:
// //	      schema:
// //	        $ref: "#/components/schemas/RentResponse"
// //	"400":
// //	  description: Bad request
// //	  content:
// //	    application/json:
// //	      schema:
// //	        $ref: "#/components/schemas/BadRequestResponse"
// //	"401":
// //	  description: Unauthorized
// //	  content:
// //	    application/json:
// //	      schema:
// //	        $ref: "#/components/schemas/UnauthorizedResponse"
// //	"500":
// //	  description: Internal server error
// //	  content:
// //	    application/json:
// //	      schema:
// //	        $ref: "#/components/schemas/ErrorResponse"
func RentBicycle(db *gorm.DB, store sessions.Store) http.HandlerFunc {
	// swagger:operation POST /bicycles/rent
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
	//   	"$ref": "#/definitions/bicycles/rent"
	// responses:
	//   '200':
	//     description: rentalss response

	return func(w http.ResponseWriter, r *http.Request) {
		// Get the user ID from the session data
		session, err := store.Get(r, "session-name")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		userID, ok := session.Values["user_id"].(uint)
		if !ok {
			http.Error(w, "Invalid session data", http.StatusBadRequest)
			return
		}

		// Parse the request body
		var rentRequest RentRequest
		err = json.NewDecoder(r.Body).Decode(&rentRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Get the bicycle from the database
		var bicycle Bicycle
		result := db.First(&bicycle, rentRequest.BicycleID)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}

		// Check if the bicycle is already rented
		if bicycle.Rented {
			http.Error(w, "Bicycle is already rented", http.StatusBadRequest)
			return
		}

		// Rent the bicycle
		bicycle.Rented = true
		result = db.Save(&bicycle)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}

		// Serialize the bicycle data to JSON format and return it to the frontend
		jsonData, err := json.Marshal(bicycle)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}

func SetupRoutes(r *mux.Router, db *gorm.DB, store sessions.Store) {
	r.HandleFunc("/bicycles/rent", RentBicycle(db, store)).Methods("POST")
}
