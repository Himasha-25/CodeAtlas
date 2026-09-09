package repository

import "time"

type Repository struct {
	ID        uint      `gorm:"primaryKey"`
	ProjectID uint      `gorm:"not null;index"`
	Name      string    `gorm:"not null"`
	StorePath string    `gorm:"not null"` // path to extracted repo on disk
	CreatedAt time.Time
}
