package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pradeep/golang-micro/authenticate"
	"github.com/pradeep/golang-micro/controller"
)

func Authenticated(router *mux.Router, authController *controller.AuthController) {
	router.HandleFunc("/users", authenticate.Authenticated(authController.GetAllEmployee)).Methods(http.MethodGet)
	router.HandleFunc("/user/{id}", authenticate.Authenticated(authController.GetEmployeeByID)).Methods(http.MethodGet)

}
