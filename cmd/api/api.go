package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pradeep/golang-micro/controller"
	auth "github.com/pradeep/golang-micro/repository/authuser"
	"github.com/pradeep/golang-micro/repository/cart"
	"github.com/pradeep/golang-micro/repository/product"
	"github.com/pradeep/golang-micro/repository/users"

	"github.com/pradeep/golang-micro/routes"
)

type ApiServer struct {
	addr string
	db   *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *ApiServer {
	return &ApiServer{addr: addr, db: db}
}
func (s *ApiServer) Run() error {

	user := users.NewBackendRepo(s.db)
	userController := controller.NewHandler(user)
	auth := auth.NewBackendRepo(s.db)
	authController := controller.NewHandular(auth)
	product := product.NewBackendRepo(s.db)
	productController := controller.NewProductHandular(product)
	cart := cart.NewBackendRepo(s.db)
	cartController := controller.NewCartHandular(cart, user)
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	routes.RegisterRoutes(subrouter, userController)
	routes.Authenticated(subrouter, authController)
	routes.ProductRoutes(subrouter, productController)
	routes.CartRoutes(subrouter, cartController)
	log.Println("lisiening port on ", s.addr)
	return http.ListenAndServe(s.addr, router)

}
