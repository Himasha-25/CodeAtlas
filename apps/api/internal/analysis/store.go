package analysis

import (
	"time"

	"gorm.io/gorm"
)

type AnalysisRun struct {
	ID           uint      `gorm:"primaryKey"`
	RepositoryID uint      `gorm:"not null;index"`
	Status       string    `gorm:"not null;default:'pending'"` // pending|running|completed|failed
	Error        string
	StartedAt    *time.Time
	CompletedAt  *time.Time
	CreatedAt    time.Time
}

type store struct{ db *gorm.DB }

func newStore(db *gorm.DB) *store { return &store{db} }

func (s *store) create(run *AnalysisRun) error { return s.db.Create(run).Error }

func (s *store) updateStatus(id uint, status, errMsg string) error {
	return s.db.Model(&AnalysisRun{}).Where("id = ?", id).
		Updates(map[string]interface{}{"status": status, "error": errMsg}).Error
}

func (s *store) findLatest(repositoryID uint) (*AnalysisRun, error) {
	var run AnalysisRun
	return &run, s.db.Where("repository_id = ?", repositoryID).
		Order("created_at desc").First(&run).Error
}
