package testutil

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewTestDB returns a gorm.DB connected to a test database.
// Set TEST_DATABASE_URL env var to use a real Postgres instance.
func NewTestDB() (*gorm.DB, error) {
	dsn := "postgres://codeatlas:codeatlas@localhost:5432/codeatlas_test?sslmode=disable"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
