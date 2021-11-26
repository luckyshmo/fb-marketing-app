package models

import (
	"time"

	"github.com/google/uuid"
)

//User model
type User struct {
	Id          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email" binding:"required"`
	Password    string    `json:"password" binding:"required"`
	PhoneNumber string    `json:"phoneNumber" db:"phone_number"`
	Amount      float64   `json:"amount" db:"amount"`
	TimeCreated time.Time `db:"time_created"`
}
