package services_test

import (
	"fmt"
	"mockery_test/internal/models"
	"mockery_test/internal/services"
	"mockery_test/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewProductService_Insert_Get(t *testing.T) {
	// Вводим тестовые значения
	testProducts := []models.Product{
		{
			ID:    "2f1afe98-63c4-4f59-bcaf-1df835602bdb",
			Name:  "mac2",
			Price: 2000000.0,
			Stock: 1,
		},
		{
			ID:    "2f1afe98-63c4-4f59-bcaf-1df835602bd0",
			Name:  "mac1",
			Price: 1000000.0,
			Stock: 10,
		},
	}

	// Мокаем репозиторий
	mockRepository := mocks.NewProductRepositoryInterface(t)

	// Создаем сервис
	service := services.NewProductService(mockRepository)

	addReq := mockRepository.EXPECT().Add(mock.AnythingOfType("models.Product"))
	getReq := mockRepository.EXPECT().Get(mock.AnythingOfType("string"))

	for i, testProduct := range testProducts {
		// Реализуем интерфейсы репозитория
		addReq.Return(nil)
		getReq.Return(testProducts[i], nil)

		// Тестируем запросы insert у репозитория
		t.Run(fmt.Sprintf("Add %s", testProduct.Name), func(t *testing.T) {
			err := service.Insert(
				testProduct.ID,
				models.InsertProductDTO{
					Name:  testProduct.Name,
					Price: testProduct.Price,
					Stock: testProduct.Stock,
				})

			assert.Nil(t, err)
		})

		// Тестируем запросы получения из репозитория
		t.Run(fmt.Sprintf("Get %s", testProduct.Name), func(t *testing.T) {
			gettedProduct, err := service.GetPrice(testProduct.ID)

			assert.Nil(t, err)
			assert.Equal(t, gettedProduct, testProduct.Price)
		})
	}
}
