package project

import "gorm.io/gorm"

type repository struct{ db *gorm.DB }

func newRepository(db *gorm.DB) *repository { return &repository{db} }

func (r *repository) create(p *Project) error { return r.db.Create(p).Error }

func (r *repository) findByUser(userID uint) ([]Project, error) {
	var projects []Project
	return projects, r.db.Where("user_id = ?", userID).Find(&projects).Error
}

func (r *repository) findOne(id, userID uint) (*Project, error) {
	var p Project
	return &p, r.db.Where("id = ? AND user_id = ?", id, userID).First(&p).Error
}

func (r *repository) update(p *Project) error { return r.db.Save(p).Error }

func (r *repository) delete(id, userID uint) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&Project{}).Error
}
