package auth

import (
	"gorm.io/gorm"
)

type repository struct{ db *gorm.DB }

func newRepository(db *gorm.DB) *repository { return &repository{db} }

func (r *repository) create(u *User) error {
	return r.db.Create(u).Error
}

func (r *repository) findByEmail(email string) (*User, error) {
	var u User
	return &u, r.db.Where("email = ?", email).First(&u).Error
}
