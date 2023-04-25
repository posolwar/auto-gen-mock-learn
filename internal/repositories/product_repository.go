package repositories

import (
	"sync"

	"mockery_test/internal/models"
)

type ProductRepository struct {
	data  map[string]models.Product
	mutex sync.RWMutex
}

func NewProductRepository() ProductRepository {
	return ProductRepository{
		data:  make(map[string]models.Product),
		mutex: sync.RWMutex{},
	}
}

func (r *ProductRepository) Add(product models.Product) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.data[product.ID] = product
	return nil
}

func (r *ProductRepository) Get(id string) (models.Product, error) {
	return r.data[id], nil
}
