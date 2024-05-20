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
