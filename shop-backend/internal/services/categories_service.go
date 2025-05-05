package services

import (
	"shop-backend/internal/domain"
	"shop-backend/internal/infra/repositories"

	"github.com/google/uuid"
)

// CategoriesService coordinates all business logic around product categories
type CategoriesService struct {
	CategoriesRepo repositories.CategoriesRepo
}

// NewCategoriesService constructs a service with its repository dependency
func NewCategoriesService(cr repositories.CategoriesRepo) *CategoriesService {
	return &CategoriesService{cr}
}

// GetAllCategoriesForAProduct returns every category belonging to a given product
// productID must be a valid UUID of an existing product
func (cs *CategoriesService) GetAllCategoriesForAProduct(productID string) ([]domain.Category, error) {
	if _, err := uuid.Parse(productID); err != nil {
		return nil, domain.Wrapf(domain.ErrInvalidUUID, "failed to parse product_id: %s", productID)
	}

	return cs.CategoriesRepo.ListAllForProduct(productID)

}
