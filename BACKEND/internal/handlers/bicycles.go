package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Bicycle struct {
	ID        uint    `json:"id"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Rented    bool    `json:"rented"`
}

func GetAllBicycles(db *gorm.DB) http.HandlerFunc {
	// swagger:operation GET /bicycles getAllBicycles
	//
	// Insert Documentation
	//
	// ---
	// produces:
	// - application/json
	// responses:
	//   '200':
	//     description: bicycle response
	//     schema:
	//       type: array
	//       items:
	//         "$ref": "#/definitions/bicycles"

	return func(w http.ResponseWriter, r *http.Request) {
		var bicycles []Bicycle
		result := db.Find(&bicycles)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}
		jsonData, err := json.Marshal(bicycles)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}

func SetupRoutes(r *mux.Router, db *gorm.DB) {
	r.HandleFunc("/bicycles", GetAllBicycles(db)).Methods("GET")
}
