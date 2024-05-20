package repository

import "github.com/pradeep/golang-micro/model"

type User interface {
	CreateEmployee(user model.User) (string, error)
	EmployeeByEmail(email string) (*model.User, error)
	EmployeeByPhone(phone string) (*model.User, error)
	GetAllProductsFromStore() ([]model.Productstore, error)
	ProductBYID(string) (*model.Productstore, error)
	UpdateProduct(string, int) error
}
