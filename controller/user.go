package controller

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pradeep/golang-micro/model"
	repositories "github.com/pradeep/golang-micro/repository"
	"github.com/pradeep/golang-micro/token"
	"github.com/pradeep/golang-micro/utils"
	"github.com/pradeep/golang-micro/validate"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	backendrepo repositories.User
}

func NewHandler(backendrepo repositories.User) *UserController {
	return &UserController{
		backendrepo: backendrepo,
	}
}
func (s *UserController) HandleRegister(w http.ResponseWriter, r *http.Request) {
	var user model.UserRegister
	if err := utils.ParseJson(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if user.Email == "" && !validate.IsValidEmail(user.Email) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("enter a valid email"))
	}
	email, err := s.backendrepo.EmployeeByEmail(user.Email)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}
	if email.Email == user.Email {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with %s already exists", email.Email))
		return
	}
	phone, err := s.backendrepo.EmployeeByPhone(user.Phone)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}
	if phone.Phone == user.Phone {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with %s already exists", phone.Phone))
		return
	}
	emp, err := s.backendrepo.CreateEmployee(model.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Phone:     user.Phone,
		Password:  user.Password,
		Email:     user.Email,
		Role:      user.Role,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, map[string]string{"user-id": emp})
}
func (s *UserController) HandleLogin(w http.ResponseWriter, r *http.Request) {
	user := new(model.UserLogin)
	if err := utils.ParseJson(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	emp, err := s.backendrepo.EmployeeByEmail(user.Email)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(emp.Password), []byte(user.Password)); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("please emter invalid credientals"))
		return
	}
	tokens, _, err := token.GenerateAllToken(emp.User_id, emp.FirstName, emp.Email, emp.Role)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	// http.SetCookie(
	// 	w, &http.Cookie{
	// 		Name:    "jwt-token",
	// 		Value:   tokens,
	// 		Expires: time.Now().Add(time.Hour * 12),
	// 	},
	// )
	utils.WriteJson(w, http.StatusOK, tokens)
	//utils.WriteJson(w, http.StatusOK, map[string]string{"message": "you are loggin successfully."})

}
func (s *UserController) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	product, err := s.backendrepo.GetAllProductsFromStore()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}
	utils.WriteJson(w, http.StatusOK, product)
}
func (s *UserController) GetProductByID(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id := param["product-id"]
	if id == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing id"))
		return
	}
	product, err := s.backendrepo.ProductBYID(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	if product.Product_ID == "" {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("no users found / invalid id"))
		return
	} else {
		utils.WriteJson(w, http.StatusOK, product)
		return
	}

}
