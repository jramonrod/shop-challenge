package services

import (
	"shop-backend/internal/domain"
	"shop-backend/internal/infra/repositories"

	"github.com/google/uuid"
)

// CartsService coordinates all business logic around the shopping cart
type CartsService struct {
	CartsRepo repositories.CartsRepo
}

// NewCartsService constructs a service with its repository dependency
func NewCartsService(cr repositories.CartsRepo) *CartsService {
	return &CartsService{cr}
}

// GetCart returns the given cart_id with the parts in the cart
func (cs *CartsService) GetCart(cartID string) (*domain.Cart, error) {
	if _, err := uuid.Parse(cartID); err != nil {
		return nil, domain.Wrapf(domain.ErrInvalidUUID, "failed to parse cart_id: %s", cartID)
	}

	return cs.CartsRepo.GetCart(cartID)
}

// CreateCart returns a cart entry with the new given parts
func (cs *CartsService) CreateCart(parts domain.PartsSelection) (*domain.Cart, error) {
	return cs.CartsRepo.CreateCart(parts)
}
