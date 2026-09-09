package repository

import "gorm.io/gorm"

type store struct{ db *gorm.DB }

func newStore(db *gorm.DB) *store { return &store{db} }

func (s *store) create(r *Repository) error { return s.db.Create(r).Error }

func (s *store) findByProject(projectID uint) ([]Repository, error) {
	var repos []Repository
	return repos, s.db.Where("project_id = ?", projectID).Find(&repos).Error
}

func (s *store) findOne(id uint) (*Repository, error) {
	var r Repository
	return &r, s.db.First(&r, id).Error
}
