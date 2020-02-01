package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"errors"

	"github.com/gorilla/mux"
	"github.com/ovidiugiorgi/wsproduct/model"
)

type ProductController struct {
	service model.ProductService
}

func NewProductController(service model.ProductService) *ProductController {
	var c = new(ProductController)
	c.service = service
	return c
}

func (c *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) error {
	var p = new(model.Product)
	err := json.NewDecoder(r.Body).Decode(p)
	if err != nil {
		return err
	}
	p, err = c.service.Add(p.ProductName, p.ProductDescription, p.VendorName)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(p)
	return err
}

func (c *ProductController) ListProducts(w http.ResponseWriter, r *http.Request) error {
	products, err := c.service.GetAll()
	if err != nil {
		return errors.New("could not fetch products")
	}
	response := make([]byte, 0, len(products))
	for _, product := range products {
		marshalled, err := json.Marshal(product)
		if err != nil {
			return errors.New("corrupted data")
		}
		response = append(response, marshalled...)
	}
	_, err = w.Write(response)
	return err
}

func (c *ProductController) GetProduct(w http.ResponseWriter, r *http.Request) error {
	productID, err := strconv.Atoi(mux.Vars(r)["productID"])
	if err != nil {
		return fmt.Errorf("invalid productID: %v", err.Error())
	}
	p, err := c.service.Get(int64(productID))
	if err != nil {
		return errors.New("product not found")
	}
	err = json.NewEncoder(w).Encode(p)
	return err
}

func (c *ProductController) RemoveProduct(w http.ResponseWriter, r *http.Request) error {
	productID, err := strconv.Atoi(mux.Vars(r)["productID"])
	if err != nil {
		return fmt.Errorf("invalid productID: %v", err.Error())
	}
	err = c.service.Remove(int64(productID))
	if err != nil {
		return fmt.Errorf("could not remove product: %v", err.Error())
	}
	w.WriteHeader(http.StatusOK)
	return err
}
