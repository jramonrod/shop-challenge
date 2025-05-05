package repositories

import (
	"shop-backend/internal/domain"
	"shop-backend/internal/infra/models"

	"gorm.io/gorm"
)

type RestrictionsRepo struct{ db *gorm.DB }

func NewRestrictionsRepo(db *gorm.DB) *RestrictionsRepo { return &RestrictionsRepo{db: db} }

// GetRestrictionsForParts returns all Restriction entries whose Principal matches any of the provided
// partIDs.
//
// The data structure is a mapping of PrincipalIDs to Restriction to  objects
func (rr *RestrictionsRepo) GetRestrictionsForParts(
	productID string,
	partIDs []string,
) (map[string][]domain.Restriction, error) {
	var restrictions []models.RestrictionModel
	err := rr.db.
		Table("restrictions AS rs").
		Joins("JOIN parts ON parts.uuid = rs.principal").
		Where("rs.principal IN ?", partIDs).
		Where("parts.product_id = ?", productID).
		Find(&restrictions).Error
	if err != nil {
		return nil, err
	}

	grouped := make(map[string][]domain.Restriction, len(restrictions))
	for _, restr := range restrictions {
		dr := restr.ToDomainObject()
		grouped[restr.Principal] = append(grouped[restr.Principal], dr)
	}

	return grouped, nil
}
