package repository

import "github.com/pradeep/golang-micro/model"

type AuthStore interface {
	GetAllEmployees() ([]model.Users, error)
	GetUserbyID(id string) (*model.Users, error)
}
