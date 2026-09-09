package dependencygraph

import "gorm.io/gorm"

type Service struct{ store *store }

func NewService(db *gorm.DB) *Service { return &Service{store: newStore(db)} }

func (s *Service) GetGraph(runID uint) ([]Dependency, error) { return s.store.listEdges(runID) }
