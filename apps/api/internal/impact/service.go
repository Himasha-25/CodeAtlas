package impact

import "gorm.io/gorm"

type Service struct{ db *gorm.DB }

func NewService(db *gorm.DB) *Service { return &Service{db} }

type ImpactResult struct {
	OriginID    uint   `json:"originId"`
	AffectedIDs []uint `json:"affectedIds"`
}

func (s *Service) Analyze(runID, symbolID uint) (*ImpactResult, error) {
	affected, err := traverse(s.db, runID, symbolID)
	if err != nil {
		return nil, err
	}
	return &ImpactResult{OriginID: symbolID, AffectedIDs: affected}, nil
}
