package model

import (
	"time"
)

type Order struct {
	Product_ID string   `json:"product-id"`
	Quantity   int      `json:"quantity"`
	Address    *Address `json:"address"`
}

type Order_Status struct {
	Order_ID    string
	User_ID     string
	Product_ID  string
	Quantity    int
	Address     *Address
	Total_Price float64
	Created_at  time.Time
}
type Address struct {
	City     string `json:"city"`
	District string `json:"district"`
	State    string `json:"state"`
	Zipcode  int    `json:"zipcode"`
}
