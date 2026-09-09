package codeexplorer

import "gorm.io/gorm"

type File struct {
	ID        uint   `gorm:"primaryKey"`
	RunID     uint   `gorm:"not null;index"`
	Path      string `gorm:"not null"`
	Language  string
	LineCount int
}

type Symbol struct {
	ID       uint   `gorm:"primaryKey"`
	FileID   uint   `gorm:"not null;index"`
	RunID    uint   `gorm:"not null;index"`
	Name     string `gorm:"not null"`
	Kind     string // function|class|interface|variable
	Line     int
	Exported bool
}

type store struct{ db *gorm.DB }

func newStore(db *gorm.DB) *store { return &store{db} }

func (s *store) listFiles(runID uint) ([]File, error) {
	var files []File
	return files, s.db.Where("run_id = ?", runID).Find(&files).Error
}

func (s *store) listSymbols(fileID uint) ([]Symbol, error) {
	var symbols []Symbol
	return symbols, s.db.Where("file_id = ?", fileID).Find(&symbols).Error
}

func (s *store) searchSymbols(runID uint, query string) ([]Symbol, error) {
	var symbols []Symbol
	return symbols, s.db.Where("run_id = ? AND name ILIKE ?", runID, "%"+query+"%").Find(&symbols).Error
}
