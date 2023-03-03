package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"BIKE-SHARING-SERVICE/internal/db"
)

type ReturnBicycleRequest struct {
	BikeID int `json:"bike_id"`
}

// // ReturnBicycle returns a bicycle rented by a user
// // @Summary Return a bicycle
// // @Description Returns a bicycle rented by a user
// // @Tags bicycles
// // @Accept json
// // @Produce json
// // @Param id path int true "Bicycle ID"
// // @Success 200 {object} Bike
// // @Failure 400 {object} ErrorResponse
// // @Router /bicycles/{id}/return [patch]
func ReturnBicycle(w http.ResponseWriter, r *http.Request) {
	// swagger:operation POST /bicycles/return
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
	//   	"$ref": "#/definitions/bicycles/return"
	// responses:
	//   '200':
	//     description: return response

	session, err := sessionStore.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Check if user is logged in
	userID, ok := session.Values["user_id"].(int)
	if !ok {
		http.Error(w, "User is not logged in", http.StatusUnauthorized)
		return
	}

	// Get bike ID from request
	var req ReturnBicycleRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get bike from DB
	bike, err := db.GetBicycle(req.BikeID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Check if bike is rented by the user
	if bike.UserID != userID {
		http.Error(w, "Bike is not rented by the user", http.StatusBadRequest)
		return
	}

	// Update bike status in DB
	err = db.UpdateBicycleStatus(req.BikeID, false, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send success response
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"success": true}`)
}
