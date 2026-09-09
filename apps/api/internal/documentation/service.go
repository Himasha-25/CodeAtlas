package documentation

import "gorm.io/gorm"

type Service struct{ store *store }

func NewService(db *gorm.DB) *Service { return &Service{store: newStore(db)} }

func (s *Service) Generate(projectID, runID uint) (*Documentation, error) {
	doc := &Documentation{
		ProjectID: projectID,
		RunID:     runID,
		Content:   generate(projectID, runID),
	}
	return doc, s.store.save(doc)
}

func (s *Service) Get(projectID uint) (*Documentation, error) {
	return s.store.findLatest(projectID)
}
