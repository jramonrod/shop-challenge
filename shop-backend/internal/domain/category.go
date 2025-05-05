package domain

// Represents a category within a product helping us structure our builder
type Category struct {
	UUID      string `json:"uuid"`
	Name      string `json:"name"`
	ProductID string `json:"product_id"`
}
