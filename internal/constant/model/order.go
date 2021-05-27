package model

import (
	uuid "github.com/satori/go.uuid"
	"time"
)
type Order struct {
	ID           uuid.UUID `json:"id"`
	Quantity     float32   `json:"quantity"`
	Unit         string    `json:"unit"`
	CertifiedBy  uuid.UUID     `json:"certified_by"`
	Customer_ID  uuid.UUID `json:"customer_Id"`
	Product_ID  uuid.UUID `json:"product_id"`
	Status       string    `json:"status"`
	ExpiredOn    time.Time `json:"expired_on"`
	CreatedOn    time.Time `json:"created_on"`
	UpdatedOn    time.Time `json:"update_on"`
}
