package model

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type User struct {
	ID        uuid.UUID  `json:"id"`
	Username  string     `json:"username,omitempty"`
	Password  string     `json:"password,omitempty"`
	FirstName string     `json:"first_name,omitempty"`
	LastName  string     `json:"last_name,omitempty"`
	Gender    string     `json:"gender,omitempty"`
	Email     string     `json:"email,omitempty"`
	Address   SubAddress `json:"address,omitempty"`
	Phone     string     `json:"phone,omitempty"`
	Orders    []Order    `json:"orders,omitempty"`
	RoleID    uuid.UUID  `json:"role_id,omitempty"`
	Role      *Role       `json:"role,omitempty"`
	CreatedOn time.Time  `json:"created_on,omitempty"`
	UpdatedOn time.Time  `json:"updated_on,omitempty"`
}
