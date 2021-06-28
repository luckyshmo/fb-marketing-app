package models

import (
	"time"

	"github.com/google/uuid"
)

//User model
type User struct {
	//'binding' is tag from GIN
	Id          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email" binding:"required"`
	Password    string    `json:"password" binding:"required"`
	PhoneNumber string    `json:"phone_number" db:"phone_number"`
	Amount      float64   `json:"amount" db:"amount"`
	DateCreated time.Time `db:"date_created"`
}
