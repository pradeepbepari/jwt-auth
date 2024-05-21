package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pradeep/golang-micro/authenticate"
	"github.com/pradeep/golang-micro/controller"
)

func CartRoutes(router *mux.Router, cartController *controller.CartController) {
	router.HandleFunc("/addtocart", authenticate.Authenticated(cartController.AddToCart)).Methods(http.MethodPost)
	router.HandleFunc("/viewcart/{cart-id}", authenticate.Authenticated(cartController.ViewCartItems)).Methods(http.MethodGet)
	router.HandleFunc("/order", authenticate.Authenticated(cartController.OrderItems)).Methods(http.MethodPost)

}
