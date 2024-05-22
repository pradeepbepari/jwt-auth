package product

import (
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/pradeep/golang-micro/config"
	"github.com/pradeep/golang-micro/model"
	"github.com/pradeep/golang-micro/query"
)

type BackendRepo struct {
	db *sql.DB
}

func NewBackendRepo(db *sql.DB) *BackendRepo {
	return &BackendRepo{db: db}
}
func (s *BackendRepo) CreateProduct(item model.Product) error {
	tx, err := s.db.Begin()
	if err != nil {
		log.Panic(err)
	}
	layout := config.Env.TimeFormat
	item.Product_ID = uuid.New()
	item.Created_at, _ = time.Parse(layout, time.Now().Format(layout))
	item.Updated_at, _ = time.Parse(layout, time.Now().Format(layout))
	_, err = tx.Exec(query.CreateProduct, item.Product_ID.String(), item.Product_Name, item.Product_Price, item.Product_Quantity, item.Created_at, item.Updated_at)
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
func (s *BackendRepo) DeleteByProductId(id string) error {
	tx, err := s.db.Begin()
	if err != nil {
		log.Panic(err)
	}
	_, err = tx.Exec(query.DeleteByid, id)
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
func (s *BackendRepo) CheckProductsByName(name string) (*model.Product, error) {
	rows, err := s.db.Query(query.GetProductByName, name)
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	defer rows.Close()
	var prod model.Product
	if rows.Next() {
		err := rows.Scan(&prod.Product_ID, &prod.Product_Name, &prod.Product_Price,
			&prod.Product_Quantity, &prod.Created_at, &prod.Updated_at)
		if err != nil {
			log.Panic(err)
			return nil, err
		}
	}
	return &prod, nil
}
func (s *BackendRepo) UpdateProductByName(id string, price float64, qty int) error {
	tx, err := s.db.Begin()
	if err != nil {
		log.Panic(err)
		return nil
	}
	layout := config.Env.TimeFormat
	time, _ := time.Parse(layout, time.Now().Format(layout))
	_, err = tx.Exec(query.UpdateProductByName, qty, price, time, id)
	if err != nil {
		log.Panic(err)
		return err
	}
	if err := tx.Commit(); err != nil {
		log.Panic(err)
		return err
	}
	return nil
}
