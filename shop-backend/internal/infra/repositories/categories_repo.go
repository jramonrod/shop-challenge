package repositories

import (
	"shop-backend/internal/domain"
	"shop-backend/internal/infra/models"
	"shop-backend/internal/mappers"

	"gorm.io/gorm"
)

type CategoriesRepo struct{ db *gorm.DB }

func NewCategoriesRepo(db *gorm.DB) *CategoriesRepo { return &CategoriesRepo{db: db} }

// Fetches all the categories for a given productID
func (r *CategoriesRepo) ListAllForProduct(productID string) ([]domain.Category, error) {
	var categories []models.CategoryModel
	if err := r.db.
		Where("product_id = ?", productID).
		Find(&categories).Error; err != nil {
		return nil, err
	}

	return mappers.ConvertAll(categories), nil
}
