package services

import (
	"shop-backend/internal/domain"
	"shop-backend/internal/infra/repositories"
)

// ProductsService provides the business logic for interacting with products
// The service interacts with the ProductsRepo to retrieve or manipulate product-related data
type ProductsService struct {
	ProductsRepo repositories.ProductsRepo
}

// NewProductsService constructs a service with its repository dependency
func NewProductsService(pr repositories.ProductsRepo) *ProductsService {
	return &ProductsService{pr}
}

// GetAllProducts retrieves all products registered in the system
// Returns a slice of Product domain objects and errors if something goes wrong
func (ps *ProductsService) GetAllProducts() ([]domain.Product, error) {
	return ps.ProductsRepo.ListAll()
}

// GetProductBySlug returns a single Product domain object and an error if the product is not
// found or if an error occurs.
func (ps *ProductsService) GetProductBySlug(slug string) (*domain.Product, error) {
	return ps.ProductsRepo.GetForSlug(slug)
}
