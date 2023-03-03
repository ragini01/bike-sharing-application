package models

import (
	"gorm.io/gorm"
)

// swagger:model
type Bicycle struct {
	gorm.Model
	// The name of the User
	// required: true
	Name      string `gorm:"not null"`
	Latitude  float64
	Longitude float64
	Rented    bool
}

func (b *Bicycle) Rent() {
	b.Rented = true
}

func (b *Bicycle) Return() {
	b.Rented = false
}
