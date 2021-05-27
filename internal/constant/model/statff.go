package model

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Staff struct {
	ID        uuid.UUID  `json:"id"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Gender    string     `json:"gender"`
	Salary    float64    `json:"salary"`
	Email     string     `json:"email"`
	Address   SubAddress `json:"address"`
	Phone     string     `json:"phone"`
	RoleID    uuid.UUID       `json:"role"`
	CreatedOn time.Time  `json:"created_on"`
	UpdatedOn time.Time  `json:"updated_on"`
}
