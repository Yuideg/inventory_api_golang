package model

import (
	uuid "github.com/satori/go.uuid"
	"time"
)
type Supplier struct {
	ID        uuid.UUID `json:"supplier_id"`
	Name      string    `json:"name"`
	Address     SubAddress `json:"address"`
	Fax       string    `json:"fax"`
	PoBox     string    `json:"po_box"`
	Email     string    `json:"email"`
	WebSite   string    `json:"web_site"`
	Phone     string    `json:"phone"`
	CreatedOn time.Time `json:"created_on"`
	UpdatedOn time.Time `json:"updated_on"`
}
