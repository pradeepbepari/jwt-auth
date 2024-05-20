package model

import (
	"time"

	"github.com/google/uuid"
)

type AddCart struct {
	Items []CartItem
}
type CartItem struct {
	Product_Id string `json:"product-id"`
	Quantity   int    `json:"quantity"`
}
type Cart struct {
	ID         uuid.UUID
	User_ID    string
	Product_Id string
	Quantity   int
	Created_at time.Time
	Updated_at time.Time
}
