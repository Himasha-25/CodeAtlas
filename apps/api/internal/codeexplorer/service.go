package codeexplorer

import "gorm.io/gorm"

type Service struct{ store *store }

func NewService(db *gorm.DB) *Service { return &Service{store: newStore(db)} }

func (s *Service) ListFiles(runID uint) ([]File, error)              { return s.store.listFiles(runID) }
func (s *Service) ListSymbols(fileID uint) ([]Symbol, error)         { return s.store.listSymbols(fileID) }
func (s *Service) Search(runID uint, q string) ([]Symbol, error)     { return s.store.searchSymbols(runID, q) }
