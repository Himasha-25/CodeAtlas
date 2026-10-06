package analysis

import (
	"time"

	"gorm.io/gorm"
)

type AnalysisRun struct {
	ID           uint       `gorm:"primaryKey"                          json:"id"`
	RepositoryID uint       `gorm:"not null;index"                      json:"repository_id"`
	Status       string     `gorm:"not null;default:'pending'"          json:"status"` // pending|running|completed|failed
	Error        string     `                                           json:"error,omitempty"`
	StartedAt    *time.Time `                                           json:"started_at,omitempty"`
	CompletedAt  *time.Time `                                           json:"completed_at,omitempty"`
	CreatedAt    time.Time  `                                           json:"created_at"`
}

type store struct{ db *gorm.DB }

func newStore(db *gorm.DB) *store { return &store{db} }

func (s *store) create(run *AnalysisRun) error { return s.db.Create(run).Error }

func (s *store) updateStatus(id uint, status, errMsg string) error {
	updates := map[string]interface{}{"status": status, "error": errMsg}
	now := time.Now()
	switch status {
	case "running":
		updates["started_at"] = &now
	case "completed", "failed":
		updates["completed_at"] = &now
	}
	return s.db.Model(&AnalysisRun{}).Where("id = ?", id).Updates(updates).Error
}

func (s *store) findLatest(repositoryID uint) (*AnalysisRun, error) {
	var run AnalysisRun
	return &run, s.db.Where("repository_id = ?", repositoryID).
		Order("created_at desc").First(&run).Error
}

func (s *store) listByRepository(repositoryID uint) ([]AnalysisRun, error) {
	var runs []AnalysisRun
	return runs, s.db.Where("repository_id = ?", repositoryID).
		Order("created_at desc").Find(&runs).Error
}
