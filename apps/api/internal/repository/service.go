package repository

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"gorm.io/gorm"
)

type Service struct{ store *store }

func NewService(db *gorm.DB) *Service { return &Service{store: newStore(db)} }

// Create registers a repository record that already exists on disk at storePath.
func (s *Service) Create(projectID uint, name, storePath string) (*Repository, error) {
	repo := &Repository{ProjectID: projectID, Name: name, StorePath: storePath}
	return repo, s.store.create(repo)
}

func (s *Service) Upload(projectID uint, name string, file multipart.File, header *multipart.FileHeader) (*Repository, error) {
	if header.Size > MaxUploadBytes {
		return nil, ErrFileTooLarge
	}

	tmpZip := filepath.Join(os.TempDir(), fmt.Sprintf("codeatlas-%d-%s", projectID, header.Filename))
	if err := saveMultipart(file, tmpZip, MaxUploadBytes); err != nil {
		return nil, fmt.Errorf("save upload: %w", err)
	}
	defer os.Remove(tmpZip)

	if err := validateZipHeader(tmpZip); err != nil {
		return nil, err
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

func (s *Service) Get(id uint) (*Repository, error) {
	return s.store.findOne(id)
}

func saveMultipart(src multipart.File, dest string, maxBytes int64) error {
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, io.LimitReader(src, maxBytes+1))
	return err
}
