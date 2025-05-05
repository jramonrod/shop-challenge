package models

import (
	"shop-backend/internal/domain"
	"time"
)

// ensure we implement the interface
var _ DomainModel[domain.Product] = &ProductModel{}

// ProductModel is the db representation for a Product object
type ProductModel struct {
	UUID      string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name      string    `gorm:"not null"`
	Slug      string    `gorm:"not null"`
	Location  string    `gorm:"default:null"`
	Status    string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (ProductModel) TableName() string {
	return "products"
}

// ToDomainObject handles converting this object to a domain object
func (pm ProductModel) ToDomainObject() domain.Product {
	return domain.Product{
		UUID:        pm.UUID,
		Name:        pm.Name,
		Location:    pm.Location,
		Slug:        pm.Slug,
		IsAvailable: pm.Status == "available",
	}
}
