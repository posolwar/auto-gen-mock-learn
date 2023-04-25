package repositories

import "mockery_test/internal/models"

//go:generate mockery --name ProductRepositoryInterface
type ProductRepositoryInterface interface {
	Add(product models.Product) error
	Get(productID string) (models.Product, error)
}
