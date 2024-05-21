package model

import (
	"time"

	"github.com/google/uuid"
)

type AddCart struct {
	Items []CartItem
}
type CartItem struct {
	Product_Id string  `json:"product-id"`
	Quantity   int     `json:"quantity"`
	Prod_Price float64 `json:"price"`
}
type Cart struct {
	ID            uuid.UUID
	User_ID       string
	Product_Id    string
	Product_Price float64
	Quantity      int
	Created_at    time.Time
	Updated_at    time.Time
}
