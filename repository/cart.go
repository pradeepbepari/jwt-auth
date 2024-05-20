package repository

import "github.com/pradeep/golang-micro/model"

type Cart interface {
	AddCartsItem(model.Cart) error
	CheckProducts(string) (*model.Cart, error)
	UpdateCart(string, string, int) (*model.Cart, error)
	PurchaseOrders(model.Order_Status) (*model.Order_Status, error)
}
