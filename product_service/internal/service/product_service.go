package service

import (
	repo "product_service/internal/repository"
)

func GetAllProducts(r repo.ProductRepository){}

func GetProductByID(r repo.ProductRepository, id string) {}

func CreateProduct(r repo.ProductRepository, product *repo.ProductRepository) {}

func UpdateProduct(r repo.ProductRepository, product *repo.ProductRepository) {}

func DeleteProduct(r repo.ProductRepository, id string) {}