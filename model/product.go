package model

import (
	"errors"
	"sync"
)

type Product struct {
	ID                 int    `json:"productId"`
	ProductName        string `json:"productName"`
	ProductDescription string `json:"productDescription"`
	VendorName         string `json:"vendorName"`
}

type ProductService interface {
	Add(productName, productDescription, vendorName string) (*Product, error)
	Get(productID int) (*Product, error)
	GetAll() ([]Product, error)
	Remove(productID int) error
}

type ProductStore struct {
	products map[int]*Product
	mu       sync.RWMutex
	id       int
}

func NewProductStore() *ProductStore {
	s := new(ProductStore)
	s.products = make(map[int]*Product)
	return s
}

func (s *ProductStore) Add(productName, productDescription, vendorName string) (*Product, error) {
	s.mu.Lock()
	s.id++
	s.mu.Unlock()

	s.mu.Lock()
	defer s.mu.Unlock()
	p := &Product{s.id, productName, productDescription, vendorName}
	s.products[s.id] = p

	return p, nil
}

func (s *ProductStore) Get(productID int) (*Product, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	p, ok := s.products[productID]
	if !ok {
		return nil, errors.New("could not find product")
	}
	return p, nil
}

func (s *ProductStore) GetAll() ([]Product, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	products := make([]Product, 0, len(s.products))
	for _, v := range s.products {
		products = append(products, *v)
	}
	return products, nil
}

func (s *ProductStore) Remove(productID int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.products[productID]; !ok {
		return errors.New("product not found")
	}

	delete(s.products, productID)
	return nil
}
