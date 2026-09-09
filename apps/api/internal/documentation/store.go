package documentation

import (
	"time"

	"gorm.io/gorm"
)

type Documentation struct {
	ID        uint      `gorm:"primaryKey"`
	ProjectID uint      `gorm:"not null;index"`
	RunID     uint      `gorm:"not null;index"`
	Content   string    `gorm:"type:text"`
	CreatedAt time.Time
}

type store struct{ db *gorm.DB }

func newStore(db *gorm.DB) *store { return &store{db} }

func (s *store) save(doc *Documentation) error { return s.db.Save(doc).Error }

func (s *store) findLatest(projectID uint) (*Documentation, error) {
	var doc Documentation
	return &doc, s.db.Where("project_id = ?", projectID).Order("created_at desc").First(&doc).Error
}
