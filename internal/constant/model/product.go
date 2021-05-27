package model

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Product struct {
	ID          uuid.UUID `json:"id"`
	Barcode     string    `json:"product_barcode"`
	ProductName string    `json:"product_name"`
	StockId     uuid.UUID      `json:"stock_id"`
	Brand       string    `json:"brand"`
	Image       string    `json:"image"`
	Cost        float32   `json:"cost"`
	Price       float32   `json:"price"`
	Unit        string    `json:"unit"`
	Discount    float32    `json:"discount"`
	Tax         float32    `json:"tax"`
	MfgDate     time.Time    `json:"mfg_date"`
	ExpiredOn   time.Time `json:"expired_on"`
	CreatedOn   time.Time `json:"created_on"`
	UpdatedOn   time.Time `json:"updated_on"`
}
