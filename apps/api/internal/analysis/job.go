package analysis

import (
	"log"
	"time"

	"github.com/codeatlas/api/analyzer"
)

func (s *Service) runJob(run *AnalysisRun, repoPath string) {
	go func() {
		now := time.Now()
		s.store.db.Model(run).Updates(map[string]interface{}{"status": "running", "started_at": &now})

		result, err := analyzer.Analyze(repoPath)
		if err != nil {
			s.store.updateStatus(run.ID, "failed", err.Error())
			log.Printf("analysis %d failed: %v", run.ID, err)
			return
		}

		if err := s.persistResult(run.ID, result); err != nil {
			s.store.updateStatus(run.ID, "failed", err.Error())
			return
		}

		done := time.Now()
		s.store.db.Model(run).Updates(map[string]interface{}{"status": "completed", "completed_at": &done})
	}()
}

func (s *Service) persistResult(runID uint, result *analyzer.Result) error {
	// TODO: persist result.Files, result.Symbols, result.Dependencies, result.Endpoints
	return nil
}
