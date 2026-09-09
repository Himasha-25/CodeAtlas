package dependencygraph

import "gorm.io/gorm"

type Dependency struct {
	ID       uint   `gorm:"primaryKey"`
	RunID    uint   `gorm:"not null;index"`
	FromID   uint   `gorm:"not null"` // file_id or symbol_id
	ToID     uint   `gorm:"not null"`
	EdgeType string `gorm:"not null"` // file|symbol|api
}

type store struct{ db *gorm.DB }

func newStore(db *gorm.DB) *store { return &store{db} }

func (s *store) listEdges(runID uint) ([]Dependency, error) {
	var deps []Dependency
	return deps, s.db.Where("run_id = ?", runID).Find(&deps).Error
}
