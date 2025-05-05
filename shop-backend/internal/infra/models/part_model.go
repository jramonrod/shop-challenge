package models

import (
	"fmt"
	"shop-backend/internal/domain"
	"time"
)

// ensure we implement the interface
var _ DomainModel[domain.Part] = PartModel{}

type PartModel struct {
	UUID       string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name       string    `gorm:"not null"`
	ProductID  string    `gorm:"not null"`
	Inventory  int       `gorm:"not null"`
	PriceCents int       `gorm:"not null"`
	CategoryID string    `gorm:"not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

func (PartModel) TableName() string {
	return "parts"
}

func (pm PartModel) ToDomainObject() domain.Part {
	return domain.Part{
		UUID:       pm.UUID,
		Name:       pm.Name,
		ProductID:  pm.ProductID,
		Inventory:  pm.Inventory,
		Price:      fmt.Sprintf("%.2f", float64(pm.PriceCents)/100.0),
		CategoryID: pm.CategoryID,
	}
}
