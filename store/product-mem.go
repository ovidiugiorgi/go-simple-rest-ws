package store

import (
	"errors"
	"sync"

	"github.com/ovidiugiorgi/wsproduct/model"
)

type InMemStore struct {
	products map[int64]*model.Product
	mu       sync.RWMutex
	id       int64
}

func NewInMemStore() *InMemStore {
	s := new(InMemStore)
	s.products = make(map[int64]*model.Product)
	return s
}

func (s *InMemStore) Add(productName, productDescription, vendorName string) (*model.Product, error) {
	s.mu.Lock()
	s.id++
	s.mu.Unlock()

	s.mu.Lock()
	defer s.mu.Unlock()
	p := &model.Product{
		ID:                 s.id,
		ProductName:        productName,
		ProductDescription: productDescription,
		VendorName:         vendorName,
	}
	s.products[s.id] = p

	return p, nil
}

func (s *InMemStore) Get(productID int64) (*model.Product, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	p, ok := s.products[productID]
	if !ok {
		return nil, errors.New("could not find product")
	}
	return p, nil
}

func (s *InMemStore) GetAll() ([]model.Product, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	products := make([]model.Product, 0, len(s.products))
	for _, v := range s.products {
		products = append(products, *v)
	}
	return products, nil
}

func (s *InMemStore) Remove(productID int64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.products[productID]; !ok {
		return errors.New("product not found")
	}

	delete(s.products, productID)
	return nil
}
