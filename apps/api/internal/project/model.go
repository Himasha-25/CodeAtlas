package project

import "time"

type Project struct {
	ID                 uint      `gorm:"primaryKey"`
	UserID             uint      `gorm:"not null;index"`
	Name               string    `gorm:"not null"`
	Description        string
	ActiveRepositoryID *uint     `gorm:"column:active_repository_id"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

type CreateRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type UpdateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProjectResponse struct {
	ID                 uint      `json:"id"`
	Name               string    `json:"name"`
	Description        string    `json:"description"`
	ActiveRepositoryID *uint     `json:"activeRepositoryId"`
	CreatedAt          time.Time `json:"createdAt"`
}
