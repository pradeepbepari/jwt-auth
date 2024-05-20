package model

import (
	"time"

	"github.com/google/uuid"
)

type ProductEntry struct {
	Product_Name     string  `json:"product_name"`
	Product_Quantity int     `json:"product_qty"`
	Product_Price    float64 `json:"product_price"`
}
type Product struct {
	Product_ID       uuid.UUID
	Product_Name     string
	Product_Quantity int
	Product_Price    float64
	Created_at       time.Time
	Updated_at       time.Time
}
type Productstore struct {
	Product_ID       string
	Product_Name     string
	Product_Quantity int
	Product_Price    float64
}
