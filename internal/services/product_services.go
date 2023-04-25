package services

import (
	"errors"

	"mockery_test/internal/models"
	"mockery_test/internal/repositories"
)

type ProductService struct {
	repo repositories.ProductRepositoryInterface
}

func NewProductService(repo repositories.ProductRepositoryInterface) ProductService {
	return ProductService{
		repo: repo,
	}
}

func (s ProductService) Insert(productID string, product models.InsertProductDTO) error {
	if len(productID) == 0 {
		return errors.New("productID can not be null")
	}

	err := s.repo.Add(models.Product{
		ID:    productID,
		Name:  product.Name,
		Price: product.Price,
		Stock: product.Stock,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s ProductService) GetPrice(productID string) (float64, error) {
	if len(productID) == 0 {
		return 0, errors.New("productID can not be null")
	}

	product, err := s.repo.Get(productID)
	if err != nil {
		return 0, err
	}

	return product.Price, nil
}
