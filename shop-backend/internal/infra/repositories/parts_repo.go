package repositories

import (
	"shop-backend/internal/domain"
	"shop-backend/internal/infra/models"
	"shop-backend/internal/mappers"

	"gorm.io/gorm"
)

// PartsRepo encapsulates all persistence operations related to Part entities.
type PartsRepo struct{ db *gorm.DB }

func NewPartRepo(db *gorm.DB) *PartsRepo { return &PartsRepo{db: db} }

// ListAllForCategory returns all domain.Part objects for a specified product and category
func (pr *PartsRepo) ListAllForCategory(productID string, categoryID string) ([]domain.Part, error) {
	var parts []models.PartModel
	if err := pr.db.
		Where("product_id = ?", productID).
		Where("category_id = ?", categoryID).
		Find(&parts).Error; err != nil {
		return nil, err
	}

	return mappers.ConvertAll(parts), nil
}

// GetCategoryMap returns a map of part UUIDs to its category UUIDs
//
// This map is useful for enforcing "only" restrictions: by knowing each selected part’s
// category, we can then ensure that within a category only the allowed part is chosen
func (pr *PartsRepo) GetCategoryMap(productID string) (map[string]string, error) {
	var parts []models.PartModel
	err := pr.db.
		Select("uuid", "category_id").
		Where("product_id = ?", productID).
		Find(&parts).
		Error
	if err != nil {
		return nil, err
	}

	result := make(map[string]string, len(parts))
	for _, part := range parts {
		result[part.UUID] = part.CategoryID
	}
	return result, nil
}

// GetNameForPart looks up the part by UUID and returns its Name. This is an auxiliary method
// that's not ideal and should be properly tackled with internationalization code, but for the
// purpose of this challenge it's fine
func (pr *PartsRepo) GetNameForPart(partUUID string) string {
	var part models.PartModel
	if err := pr.db.
		Select("name").
		Where("uuid = ?", partUUID).
		First(&part).Error; err != nil {
		return ""
	}

	return part.Name
}
