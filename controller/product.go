package controller

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pradeep/golang-micro/model"
	"github.com/pradeep/golang-micro/repository"
	"github.com/pradeep/golang-micro/utils"
)

type ProductController struct {
	product repository.Product
}

func NewProductHandular(product repository.Product) *ProductController {
	return &ProductController{product: product}
}

func (s *ProductController) AddProduct(w http.ResponseWriter, r *http.Request) {
	role := r.Header.Get("role")
	if role == "ADMIN" || role == "Admin" || role == "admin" {
		var product model.ProductEntry
		if err := utils.ParseJson(r, &product); err != nil {
			utils.WriteError(w, http.StatusBadRequest, err)
			return
		}
		prod, err := s.product.CheckProductsByName(product.Product_Name)
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("internal error"))
			return
		}
		if prod.Product_Name == product.Product_Name {
			prod.Product_Price = product.Product_Price
			prod.Product_Quantity += product.Product_Quantity
			err := s.product.UpdateProductByName(prod.Product_ID.String(), prod.Product_Price, prod.Product_Quantity)
			if err != nil {
				utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("internal error"))
				return
			}
			utils.WriteJson(w, http.StatusOK, map[string]string{"message": "product added"})
			return
		}
		ok := s.product.CreateProduct(model.Product{
			Product_Name:     product.Product_Name,
			Product_Quantity: product.Product_Quantity,
			Product_Price:    product.Product_Price,
		})
		if ok != nil {
			utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("internal error"))
			return
		}
		utils.WriteJson(w, http.StatusOK, map[string]string{"message": "product added"})
		return

	}
	utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("you are not authorised to do this"))
}
func (s *ProductController) DeleteProductById(w http.ResponseWriter, r *http.Request) {
	role := r.Header.Get("role")
	if role == "ADMIN" || role == "Admin" || role == "admin" {
		param := mux.Vars(r)
		id := param["prod-id"]
		if err := s.product.DeleteByProductId(id); err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}
		utils.WriteJson(w, http.StatusOK, map[string]string{"message": "product deleted"})
		return
	}
	utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("you are not authorised to do this"))
}
