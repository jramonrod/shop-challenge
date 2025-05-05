package services

import (
	"shop-backend/internal/domain"
	"shop-backend/internal/infra/repositories"

	"github.com/google/uuid"
)

// PartsService coordinates all business logic around product parts
type PartsService struct {
	PartsRepo repositories.PartsRepo
}

// NewPartsService constructs a service with its repository dependency
func NewPartsService(pr repositories.PartsRepo) *PartsService {
	return &PartsService{pr}
}

// GetAllPartsInCategory returns every part belonging to a given category and product
// productID must be a valid UUID of an existing product
func (ps *PartsService) GetAllPartsInCategory(productID string, categoryID string) ([]domain.Part, error) {
	if _, err := uuid.Parse(productID); err != nil {
		return nil, domain.Wrapf(domain.ErrInvalidUUID, "failed to parse product_id: %s", productID)
	}

	if _, err := uuid.Parse(categoryID); err != nil {
		return nil, domain.Wrapf(domain.ErrInvalidUUID, "failed to parse category_id: %s", categoryID)
	}

	return ps.PartsRepo.ListAllForCategory(productID, categoryID)
}
