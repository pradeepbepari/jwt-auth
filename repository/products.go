package repository

import "github.com/pradeep/golang-micro/model"

type Product interface {
	CreateProduct(model.Product) error
	DeleteByProductId(string) error
}
