package model

import (
	uuid "github.com/satori/go.uuid"
	"time"
)
type Stock struct {
	ID           uuid.UUID  `json:"id"`
	ProductsCode  string  `json:"products_code"`
	ProductsName string     `json:"products_name"`
	No_products     int        `json:"capacity"`
	SupplierID   uuid.UUID `json:"supplier_id"`
	Supplier_List   []Supplier `json:"supplier_list"`
	Address     SubAddress `json:"address"`
	CreatedOn    time.Time  `json:"created_on"`
	UpdatedOn    time.Time  `json:"updated_on"`
}