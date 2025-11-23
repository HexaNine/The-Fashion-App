package repository

import (
	e "product_service/internal/entity"

	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	collection mongo.Collection
}

type ProductRepositoryInterface interface {
	GetAllProducts(repo *ProductRepository) ([]e.Product, error)
	GetProductByID(repo *ProductRepository, id string) (*e.Product, error)
	CreateProduct(repo *ProductRepository, product *e.Product) error
	UpdateProduct(repo *ProductRepository, product *e.Product) error
	DeleteProduct(repo *ProductRepository, id string) error
}

