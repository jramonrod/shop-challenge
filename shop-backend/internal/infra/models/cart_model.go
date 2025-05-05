package models

import (
	"shop-backend/internal/domain"
	"time"
)

// ensure we implement the interface
var _ DomainModel[domain.Cart] = &CartModel{}

// RestrictionModel is the db representation for a Product object
type CartModel struct {
	UUID      string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Status    string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (CartModel) TableName() string {
	return "carts"
}

// ToDomainObject handles converting this object to a domain object
//
// Note the resulting domain object from this method requires joins from multiple tables
// to enrich this please use the repository method
func (cm CartModel) ToDomainObject() domain.Cart {
	return domain.Cart{
		UUID:     cm.UUID,
		Status:   cm.Status,
		Products: map[string][]domain.Part{},
	}
}
