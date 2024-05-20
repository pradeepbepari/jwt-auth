package users

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/pradeep/golang-micro/config"
	"github.com/pradeep/golang-micro/model"
	"github.com/pradeep/golang-micro/query"
	"golang.org/x/crypto/bcrypt"
)

type BackendRepo struct {
	db *sql.DB
}

func NewBackendRepo(db *sql.DB) *BackendRepo {
	return &BackendRepo{db: db}
}
func (s *BackendRepo) EmployeeByEmail(email string) (*model.User, error) {
	rows, err := s.db.Query(query.GetEmail, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var user model.User
	if rows.Next() {
		err := rows.Scan(&user.Uuid, &user.FirstName, &user.LastName, &user.Password, &user.Email,
			&user.Phone, &user.Role, &user.User_id,
			&user.Created_at, &user.Updated_at)
		if err != nil {
			return nil, err
		}
	}
	return &user, nil
}
func (s *BackendRepo) EmployeeByPhone(phone string) (*model.User, error) {
	rows, err := s.db.Query(query.GetPhone, phone)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var user model.User
	if rows.Next() {
		err := rows.Scan(&user.Uuid, &user.FirstName, &user.LastName, &user.Password, &user.Email,
			&user.Phone, &user.Role, &user.User_id,
			&user.Created_at, &user.Updated_at)
		if err != nil {
			return nil, err
		}
	}
	return &user, nil
}
func (s *BackendRepo) CreateEmployee(user model.User) (string, error) {
	tx, err := s.db.Begin()
	if err != nil {
		log.Panic(err)
	}
	user.Uuid = uuid.New()
	user.User_id = user.Uuid.String()
	layout := config.Env.TimeFormat
	user.Created_at, _ = time.Parse(layout, time.Now().Format(layout))
	user.Updated_at, _ = time.Parse(layout, time.Now().Format(layout))
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		log.Fatal(err)
	}
	_, err = tx.Exec(query.CreateUser, user.Uuid.String(), user.FirstName, user.LastName, pass, user.Email, user.Phone, user.Role, user.User_id, user.Created_at, user.Updated_at)
	if err != nil {
		fmt.Printf("error while executing query")
		log.Panic(err)
		return "", err
	}
	err = tx.Commit()
	if err != nil {
		fmt.Printf("commit failed")
		log.Panic(err)
	}
	return user.Uuid.String(), nil
}
func (s *BackendRepo) GetAllProductsFromStore() ([]model.Productstore, error) {
	rows, err := s.db.Query(query.GetAllProducts)
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	defer rows.Close()
	var products []model.Product
	for rows.Next() {
		var product model.Product
		err := rows.Scan(&product.Product_ID, &product.Product_Name,
			&product.Product_Price, &product.Product_Quantity,
			&product.Created_at, &product.Updated_at)
		if err != nil {
			log.Panic(err)
			return nil, err
		}
		products = append(products, product)
	}
	productlist := make([]model.Productstore, len(products))
	for i, product := range products {
		productlist[i] = model.Productstore{
			Product_ID:       product.Product_ID.String(),
			Product_Name:     product.Product_Name,
			Product_Price:    product.Product_Price,
			Product_Quantity: product.Product_Quantity,
		}
	}
	return productlist, nil
}
func (s *BackendRepo) ProductBYID(id string) (*model.Productstore, error) {
	rows, err := s.db.Query(query.GetProductByID, id)
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	defer rows.Close()
	var prod model.Product
	if rows.Next() {
		err := rows.Scan(&prod.Product_ID, &prod.Product_Name,
			&prod.Product_Price, &prod.Product_Quantity,
			&prod.Created_at, &prod.Updated_at)
		if err != nil {
			log.Panic(err)
			return nil, err
		}
	}
	return &model.Productstore{
		Product_ID:       prod.Product_ID.String(),
		Product_Name:     prod.Product_Name,
		Product_Price:    prod.Product_Price,
		Product_Quantity: prod.Product_Quantity,
	}, nil
}
func (s *BackendRepo) UpdateProduct(id string, qty int) error {
	tx, err := s.db.Begin()
	if err != nil {
		log.Panic(err)
	}
	layout := config.Env.TimeFormat
	Updated_at, _ := time.Parse(layout, time.Now().Format(layout))
	_, err = tx.Exec(query.UpdateProduct, qty, Updated_at, id)
	if err != nil {
		log.Println("error while executing query")
		log.Panic(err)
		return err
	}
	err = tx.Commit()
	if err != nil {
		log.Println("commit failed")
		log.Panic(err)
	}
	return nil
}
