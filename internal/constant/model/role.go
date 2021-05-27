package model

import uuid "github.com/satori/go.uuid"

type Role struct {
	ID uuid.UUID    `json:"id"`
	Name   string `json:"name"`
}
