package model

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Customer struct {
	ID        uuid.UUID  `json:"id"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Gender    string `json:"gender"`
	Email     string     `json:"email"`
	Address    SubAddress `json:"address"`
	Phone     string     `json:"phone"`
	Orders    []OrderJson   `json:"orders"`
	RoleID    uuid.UUID     `json:"role_id"`
	CreatedOn time.Time  `json:"created_on"`
	UpdatedOn time.Time  `json:"updated_on"`
}
