package models

import "github.com/google/uuid"

//User model
type User struct {
	//'binding' is tag from GIN
	Id          uuid.UUID `json:"-" db:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email" binding:"required"`
	Password    string    `json:"password" binding:"required"`
	PhoneNumber string    `json:"phoneNumber"`
	Amount      float64   `json:"amount" db:"amount"`
}
