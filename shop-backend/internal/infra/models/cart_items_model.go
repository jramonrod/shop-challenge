package models

import "time"

// CartItemsModel is the db representation for an item in a cart
//
// Note unlike the other models, this does not refer to a domain object directly so it does
// not implement the DomainModel interface
type CartItemsModel struct {
	CartID    string    `gorm:"not null"`
	PartID    string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func (CartItemsModel) TableName() string {
	return "cart_items"
}
