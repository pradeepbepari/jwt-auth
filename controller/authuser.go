package controller

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	repositories "github.com/pradeep/golang-micro/repository"
	"github.com/pradeep/golang-micro/utils"
)

type AuthController struct {
	auth repositories.AuthStore
}

func NewHandular(auth repositories.AuthStore) *AuthController {
	return &AuthController{auth: auth}
}
func (s *AuthController) GetAllEmployee(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("id")
	emp, err := s.auth.GetUserbyID(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("internal server error"))
	}
	if emp.Role == "ADMIN" || emp.Role == "Admin" || emp.Role == "admin" {
		employees, err := s.auth.GetAllEmployees()
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}
		utils.WriteJson(w, http.StatusOK, employees)
		return
	}
	utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("you are not authorised to do this"))
}

func (s *AuthController) GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	role := r.Header.Get("role")
	param := mux.Vars(r)
	id := param["id"]
	if role == "ADMIN" || role == "Admin" || role == "admin" {
		emp, err := s.auth.GetUserbyID(id)
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}
		if emp.User_id == "" {
			utils.WriteError(w, http.StatusNotFound, fmt.Errorf("no users found / invalid id"))
			return
		} else {
			utils.WriteJson(w, http.StatusOK, emp)
			return
		}

	}
	utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("you are not authorised to do this"))
}
