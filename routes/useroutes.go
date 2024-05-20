package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pradeep/golang-micro/controller"
)

func RegisterRoutes(router *mux.Router, userController *controller.UserController) {
	router.HandleFunc("/register", userController.HandleRegister).Methods(http.MethodPost)
	router.HandleFunc("/login", userController.HandleLogin).Methods(http.MethodPost)
	router.HandleFunc("/products", userController.GetAllProducts).Methods(http.MethodGet)
	router.HandleFunc("/product/{product-id}", userController.GetProductByID).Methods(http.MethodGet)

}
