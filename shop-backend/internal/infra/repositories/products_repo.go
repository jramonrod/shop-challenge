package repositories

import (
	"errors"
	"shop-backend/internal/domain"
	"shop-backend/internal/infra/models"
	"shop-backend/internal/mappers"

	"gorm.io/gorm"
)

type ProductsRepo struct{ db *gorm.DB }

func NewProductsRepo(db *gorm.DB) *ProductsRepo { return &ProductsRepo{db: db} }

// ListAll lists all the products in the application
func (pr *ProductsRepo) ListAll() ([]domain.Product, error) {
	var products []models.ProductModel
	if err := pr.db.Find(&products).Error; err != nil {
		return nil, err
	}

	return mappers.ConvertAll(products), nil
}

// GetForSlug is a useful utility to help the user user the product slug (which is unique) to look for
// product rows.
func (pr *ProductsRepo) GetForSlug(slug string) (*domain.Product, error) {
	var product models.ProductModel

	if err := pr.db.Where("slug = ?", slug).First(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.Wrapf(domain.ErrProductNotFound, "failed to find product: %s", slug)
		}

		return nil, err
	}

	var result = product.ToDomainObject()
	return &result, nil
}
