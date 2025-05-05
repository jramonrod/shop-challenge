package domain

// Product represents a top-level domain entity in the system.
//
// Although each product includes a 'slug' (a URL-safe, human-readable, and unique identifier),
// We retain a system-generated 'UUID' as the canonical primary key for consistency across domain models.
// This approach protects against future requirements such as slug renaming, distributed data systems, or
// cross-domain entity references where a UUID is preferred for stability.
type Product struct {
	UUID        string `json:"uuid"`
	Name        string `json:"name"`
	Location    string `json:"location"`
	IsAvailable bool   `json:"is_available"`
	Slug        string `json:"slug"`
}
