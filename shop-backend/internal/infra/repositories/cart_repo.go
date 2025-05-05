package repositories

import (
	"fmt"
	"shop-backend/internal/domain"
	"shop-backend/internal/infra/models"

	"gorm.io/gorm"
)

type CartsRepo struct{ db *gorm.DB }

func NewCartsRepo(db *gorm.DB) *CartsRepo { return &CartsRepo{db: db} }

// GetCart fetches a cart and enriches it to contain the product those parts belong to
// as well as the part information.
func (cr *CartsRepo) GetCart(cartID string) (*domain.Cart, error) {
	type row struct {
		Slug string `gorm:"column:slug"`
		models.PartModel
	}

	rows := []row{}
	err := cr.db.
		Table("carts").
		Select(`
			products.slug 		AS slug,
			parts.uuid       	AS uuid,
			parts.name       	AS name,
			parts.product_id 	AS product_id,
			parts.inventory  	AS inventory,
			parts.price_cents 	AS price_cents,
			parts.category_id 	AS category_id,
			parts.created_at  	AS created_at,
			parts.updated_at  	AS updated_at
		`).
		Joins("JOIN cart_items ON carts.uuid = cart_items.cart_id").
		Joins("JOIN parts ON parts.uuid = cart_items.part_id").
		Joins("JOIN products ON parts.product_id = products.uuid").
		Where("carts.uuid = ?", cartID).
		Where("products.status = ?", "available").
		Find(&rows).Error
	if err != nil {
		return nil, err
	}

	result := domain.Cart{
		UUID:     cartID,
		Products: map[string][]domain.Part{},
	}
	for _, r := range rows {
		result.Products[r.Slug] = append(result.Products[r.Slug], domain.Part{
			UUID:       r.UUID,
			Name:       r.Name,
			Inventory:  r.Inventory,
			Price:      fmt.Sprintf("%.2f", float64(r.PriceCents)/100.0),
			ProductID:  r.ProductID,
			CategoryID: r.CategoryID,
		})
	}

	return &result, nil
}

// Creates an empty cart and relates its cart_items table with all the part uuids
func (cr *CartsRepo) CreateCart(parts domain.PartsSelection) (*domain.Cart, error) {
	var newCart models.CartModel = models.CartModel{Status: "active"}
	err := cr.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&newCart).Error; err != nil {
			return fmt.Errorf("insert cart: %w", err)
		}

		items := make([]models.CartItemsModel, len(parts))
		for i, partID := range parts {
			items[i] = models.CartItemsModel{
				CartID: newCart.UUID,
				PartID: partID,
			}
		}
		if len(items) > 0 {
			if err := tx.Create(&items).Error; err != nil {
				return fmt.Errorf("insert cart_items: %w", err)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return cr.GetCart(newCart.UUID)
}
