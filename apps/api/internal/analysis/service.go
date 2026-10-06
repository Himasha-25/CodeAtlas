package analysis

import (
	"fmt"

	"github.com/codeatlas/api/internal/repository"
	"gorm.io/gorm"
)

type Service struct {
	store   *store
	repoSvc *repository.Service
}

func NewService(db *gorm.DB, repoSvc *repository.Service) *Service {
	return &Service{store: newStore(db), repoSvc: repoSvc}
}

func (s *Service) Start(repositoryID uint) (*AnalysisRun, error) {
	repo, err := s.repoSvc.Get(repositoryID)
	if err != nil {
		return nil, fmt.Errorf("repository not found: %w", err)
	}
	run := &AnalysisRun{RepositoryID: repositoryID, Status: "pending"}
	if err := s.store.create(run); err != nil {
		return nil, err
	}
	s.runJob(run, repo.StorePath)
	return run, nil
}

func (s *Service) Status(repositoryID uint) (*AnalysisRun, error) {
	return s.store.findLatest(repositoryID)
}

func (s *Service) History(repositoryID uint) ([]AnalysisRun, error) {
	return s.store.listByRepository(repositoryID)
}
