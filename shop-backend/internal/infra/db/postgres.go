// infra/db.go
package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// MustConnectDB initializes a GORM DB connection or panics.
func MustConnectDB(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}

	return db
}
