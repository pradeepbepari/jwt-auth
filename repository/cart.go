package repository

import "github.com/pradeep/golang-micro/model"

type Cart interface {
	AddCartsItem(model.Cart) (*model.Cart, error)
	CheckProductsInCart(string) (*model.Cart, error)
	UpdateCart(string, string, int) (*model.Cart, error)
	PurchaseOrders(model.Order_Status) (*model.Order_Status, error)
	ViewCart(string) (*model.Cart, error)
}
