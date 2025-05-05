package domain

// Part is an entity belonging to a category that is specific to that category in the product
// of which the category belongs to.
//
// I debated the idea of independent Parts irregardless of the product, but considering how we want
// to expand in the future conflicts could rise up, such as wheels for bikes, and wheels for roller blades.
// This current approach handles that just fine because the products would be different.
//
// Note price is a string to avoid leaky abstractions with decimals
type Part struct {
	UUID       string `json:"uuid"`
	Name       string `json:"name"`
	Inventory  int    `json:"inventory"`
	Price      string `json:"price"`
	ProductID  string `json:"product_id"`
	CategoryID string `json:"category_id"`
}
