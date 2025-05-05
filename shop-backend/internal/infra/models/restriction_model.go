package models

import (
	"shop-backend/internal/domain"
)

// ensure we implement the interface
var _ DomainModel[domain.Restriction] = &RestrictionModel{}

// RestrictionModel is the db representation for a Product object
type RestrictionModel struct {
	UUID      string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Principal string `gorm:"not null"`
	Resource  string `gorm:"not null"`
	Effect    string `gorm:"not null"`
}

func (RestrictionModel) TableName() string {
	return "restrictions"
}

// ToDomainObject handles converting this object to a domain object
func (rm RestrictionModel) ToDomainObject() domain.Restriction {
	effect, err := domain.ParseRestrictionEffect(rm.Effect)
	if err != nil {
		panic("model cannot be turned to domain object, catastrophic error exit")
	}
	return domain.Restriction{
		UUID:        rm.UUID,
		PrincipalID: rm.Principal,
		ResourceID:  rm.Resource,
		Effect:      effect,
	}
}
