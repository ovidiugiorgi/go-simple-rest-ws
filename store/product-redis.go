package store

import (
	"encoding/json"
	"errors"
	"strconv"
	"sync"

	"github.com/ovidiugiorgi/wsproduct/model"
)

const ID_KEY = "product-counter"

type RedisStore struct {
	mu    sync.RWMutex
	redis *Redis
}

func NewRedisStore() *RedisStore {
	s := new(RedisStore)
	s.redis = new(Redis)
	return s
}

func (s *RedisStore) Add(productName, productDescription, vendorName string) (*model.Product, error) {
	id, err := s.getNextID()
	if err != nil {
		return nil, err
	}

	p := &model.Product{
		ID:                 id,
		ProductName:        productName,
		ProductDescription: productDescription,
		VendorName:         vendorName,
	}
	v, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	key := s.getKey(p.ID)
	err = s.redis.Client().Set(key, v, 0).Err()
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (s *RedisStore) Get(productID int64) (*model.Product, error) {
	key := s.getKey(productID)
	res, err := s.redis.Client().Get(key).Result()
	if err != nil {
		return nil, errors.New("could not find product")
	}
	var p = new(model.Product)
	err = json.Unmarshal([]byte(res), p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (s *RedisStore) GetAll() ([]model.Product, error) {
	count := s.getID()
	products := make([]model.Product, 0, count)
	var i int64
	// TODO: Use goroutines
	for i = 1; i <= count; i++ {
		p, err := s.Get(i)
		if err != nil {
			continue
		}
		products = append(products, *p)
	}
	return products, nil
}

func (s *RedisStore) Remove(productID int64) error {
	d, err := s.redis.Client().Del(s.getKey(productID)).Result()
	if err != nil {
		return err
	}
	if d == 0 {
		return errors.New("product not found")
	}
	return nil
}

func (s *RedisStore) getNextID() (int64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	id, err := s.redis.Client().Incr(ID_KEY).Result()
	if err != nil {
		return 0, errors.New("could not generate ID")
	}
	return id, nil
}

func (s *RedisStore) getID() int64 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	stringID, err := s.redis.Client().Get(ID_KEY).Result()
	if err != nil {
		return 0
	}
	id, err := strconv.Atoi(stringID)
	if err != nil {
		return 0
	}
	return int64(id)
}

func (s *RedisStore) getKey(productID int64) string {
	return strconv.Itoa(int(productID))
}
