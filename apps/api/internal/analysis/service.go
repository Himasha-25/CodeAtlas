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
	repos, err := s.repoSvc.List(0) // caller should pass projectID; simplified here
	_ = repos
	if err != nil {
		return nil, fmt.Errorf("load repository: %w", err)
	}
	run := &AnalysisRun{RepositoryID: repositoryID, Status: "pending"}
	if err := s.store.create(run); err != nil {
		return nil, err
	}
	// TODO: look up repo.StorePath from repository store
	s.runJob(run, "")
	return run, nil
}

func (s *Service) Status(repositoryID uint) (*AnalysisRun, error) {
	return s.store.findLatest(repositoryID)
}
