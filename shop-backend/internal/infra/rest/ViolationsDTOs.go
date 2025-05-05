package rest

import "shop-backend/internal/domain"

type ViolationsDTO struct {
	IsValid    bool               `json:"is_valid"`
	Violations []domain.Violation `json:"violations"`
}
