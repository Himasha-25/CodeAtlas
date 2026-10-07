package project

import "gorm.io/gorm"

type Service struct{ repo *repository }

func NewService(db *gorm.DB) *Service { return &Service{repo: newRepository(db)} }

func (s *Service) Create(userID uint, req CreateRequest) (*Project, error) {
	p := &Project{UserID: userID, Name: req.Name, Description: req.Description}
	return p, s.repo.create(p)
}

func (s *Service) List(userID uint) ([]Project, error) { return s.repo.findByUser(userID) }

func (s *Service) Get(id, userID uint) (*Project, error) { return s.repo.findOne(id, userID) }

func (s *Service) Update(id, userID uint, req UpdateRequest) (*Project, error) {
	p, err := s.repo.findOne(id, userID)
	if err != nil {
		return nil, err
	}
	if req.Name != "" {
		p.Name = req.Name
	}
	if req.Description != "" {
		p.Description = req.Description
	}
	return p, s.repo.update(p)
}

func (s *Service) Delete(id, userID uint) error { return s.repo.delete(id, userID) }

// SetActiveRepository sets or clears (repoID == nil) the active repository for a project.
func (s *Service) SetActiveRepository(projectID, userID uint, repoID *uint) error {
	return s.repo.setActiveRepo(projectID, userID, repoID)
}
