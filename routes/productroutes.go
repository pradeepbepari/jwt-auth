package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pradeep/golang-micro/authenticate"
	"github.com/pradeep/golang-micro/controller"
)

func ProductRoutes(router *mux.Router, productController *controller.ProductController) {
	router.HandleFunc("/addproduct", authenticate.Authenticated(productController.AddProduct)).Methods(http.MethodPost)
	router.HandleFunc("/delete/{prod-id}", authenticate.Authenticated(productController.DeleteProductById)).Methods(http.MethodDelete)

}
