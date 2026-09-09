package repository

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"

	"gorm.io/gorm"
)

type Service struct{ store *store }

func NewService(db *gorm.DB) *Service { return &Service{store: newStore(db)} }

func (s *Service) Upload(projectID uint, name string, file multipart.File, header *multipart.FileHeader) (*Repository, error) {
	tmpZip := filepath.Join(os.TempDir(), header.Filename)
	if err := saveMultipart(file, tmpZip); err != nil {
		return nil, fmt.Errorf("save upload: %w", err)
	}
	destDir := filepath.Join("data", "repos", fmt.Sprintf("%d", projectID), name)
	if err := extractZip(tmpZip, destDir); err != nil {
		return nil, fmt.Errorf("extract zip: %w", err)
	}
	repo := &Repository{ProjectID: projectID, Name: name, StorePath: destDir}
	return repo, s.store.create(repo)
}

func (s *Service) List(projectID uint) ([]Repository, error) {
	return s.store.findByProject(projectID)
}

func saveMultipart(src multipart.File, dest string) error {
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()
	buf := make([]byte, 32*1024)
	for {
		n, err := src.Read(buf)
		if n > 0 {
			out.Write(buf[:n])
		}
		if err != nil {
			break
		}
	}
	return nil
}
