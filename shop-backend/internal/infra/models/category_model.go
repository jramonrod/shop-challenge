package models

import (
	"shop-backend/internal/domain"
	"time"
)

// ensure we implement the interface
var _ DomainModel[domain.Category] = CategoryModel{}

type CategoryModel struct {
	UUID      string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name      string    `gorm:"not null"`
	ProductID string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (CategoryModel) TableName() string {
	return "categories"
}

func (cg CategoryModel) ToDomainObject() domain.Category {
	return domain.Category{
		UUID:      cg.UUID,
		Name:      cg.Name,
		ProductID: cg.ProductID,
	}
}
