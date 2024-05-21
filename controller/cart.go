package controller

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pradeep/golang-micro/model"
	repository "github.com/pradeep/golang-micro/repository"
	"github.com/pradeep/golang-micro/utils"
)

type CartController struct {
	cart    repository.Cart
	product repository.User
}

func NewCartHandular(cart repository.Cart, product repository.User) *CartController {
	return &CartController{cart: cart, product: product}
}
func (s *CartController) AddToCart(w http.ResponseWriter, r *http.Request) {
	user_id := r.Header.Get("id")
	var cart model.CartItem
	if err := utils.ParseJson(r, &cart); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	prod, err := s.product.ProductBYID(cart.Product_Id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	product, err := s.cart.CheckProductsInCart(cart.Product_Id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if product.Product_Id == cart.Product_Id {
		product.Quantity += cart.Quantity
		_, err := s.cart.UpdateCart(product.ID.String(), product.Product_Id, product.Quantity)
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}
	} else {
		cart, err := s.cart.AddCartsItem(model.Cart{
			User_ID:       user_id,
			Product_Id:    prod.Product_ID,
			Quantity:      cart.Quantity,
			Product_Price: prod.Product_Price,
		})
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}
		utils.WriteJson(w, http.StatusOK, map[string]string{"cart-id": cart.ID.String()})
	}

}
func (s *CartController) ViewCartItems(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	cart_id := param["cart-id"]
	cart, err := s.cart.ViewCart(cart_id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJson(w, http.StatusOK, cart)
}
func (s *CartController) OrderItems(w http.ResponseWriter, r *http.Request) {
	user_id := r.Header.Get("id")
	var order model.Order
	if err := utils.ParseJson(r, &order); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	product, err := s.product.ProductBYID(order.Product_ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	if product.Product_Quantity < order.Quantity {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("insufficient quantity"))
		return
	}
	total_price := product.Product_Price * float64(order.Quantity)
	orders, err := s.cart.PurchaseOrders(
		model.Order_Status{
			User_ID:     user_id,
			Product_ID:  product.Product_ID,
			Quantity:    order.Quantity,
			Address:     order.Address,
			Total_Price: total_price,
		},
	)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	qty := product.Product_Quantity - order.Quantity
	err = s.product.UpdateProduct(product.Product_ID, qty)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJson(w, http.StatusOK, map[string]string{"order-id": orders.Order_ID})
}
