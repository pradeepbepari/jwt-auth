package cart

import (
	"database/sql"
	"fmt"
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
func (s *BackendRepo) AddCartsItem(item model.Cart) (*model.Cart, error) {
	tx, err := s.db.Begin()
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	item.ID = uuid.New()
	layout := config.Env.TimeFormat
	item.Created_at, _ = time.Parse(layout, time.Now().Format(layout))
	item.Updated_at, _ = time.Parse(layout, time.Now().Format(layout))
	_, err = tx.Exec(query.AddCart, item.ID.String(), item.User_ID, item.Product_Id, item.Product_Price, item.Quantity, item.Created_at, item.Updated_at)
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		log.Panic(err)
		return nil, err
	}
	return &model.Cart{
		ID: item.ID,
	}, err
}
func (s *BackendRepo) CheckProductsInCart(pid string) (*model.Cart, error) {
	log.Println(pid)
	rows, err := s.db.Query(query.CheckProductInCart, pid)
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	defer rows.Close()
	var cart model.Cart
	if rows.Next() {
		err := rows.Scan(&cart.ID, &cart.User_ID, &cart.Product_Id, &cart.Product_Price, &cart.Quantity, &cart.Created_at, &cart.Updated_at)
		if err != nil {
			log.Panic(err)
			return nil, err
		}
	}
	log.Println(cart)
	return &model.Cart{
		ID:            cart.ID,
		User_ID:       cart.User_ID,
		Product_Id:    cart.Product_Id,
		Product_Price: cart.Product_Price,
		Quantity:      cart.Quantity,
		Created_at:    cart.Created_at,
		Updated_at:    cart.Updated_at,
	}, nil
}
func (s *BackendRepo) UpdateCart(id, cart_id string, qty int) (*model.Cart, error) {
	tx, err := s.db.Begin()
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	var cart model.Cart
	layout := config.Env.TimeFormat
	cart.Updated_at, _ = time.Parse(layout, time.Now().Format(layout))
	_, err = tx.Exec(query.UpdateCart, qty, cart.Updated_at, id, cart_id)
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		log.Panic(err)
		return nil, err
	}
	return nil, nil
}
func (s *BackendRepo) PurchaseOrders(order model.Order_Status) (*model.Order_Status, error) {
	tx, err := s.db.Begin()
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	order.Order_ID = uuid.New().String()
	address := fmt.Sprintf("%s,%s,%s,%d", order.Address.City, order.Address.District, order.Address.State, order.Address.Zipcode)
	layout := config.Env.TimeFormat
	order.Created_at, _ = time.Parse(layout, time.Now().Format(layout))
	_, err = tx.Exec(query.Orderitem, order.Order_ID, order.User_ID,
		order.Product_ID, order.Quantity, address, order.Total_Price, order.Created_at)
	if err != nil {
		log.Panic(err)
		return nil, nil
	}
	if err := tx.Commit(); err != nil {
		log.Panic(err)
		return nil, nil
	}
	return &model.Order_Status{
		Order_ID: order.Order_ID,
	}, nil
}

func (s *BackendRepo) ViewCart(cartId string) (*model.Cart, error) {
	rows, err := s.db.Query(query.ViewCart, cartId)
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	defer rows.Close()
	var cart model.Cart
	if rows.Next() {
		err := rows.Scan(&cart.ID, &cart.User_ID, &cart.Product_Id, &cart.Product_Price, &cart.Quantity, &cart.Created_at, &cart.Updated_at)
		if err != nil {
			log.Panic(err)
			return nil, err
		}
	}
	return &cart, nil
}
