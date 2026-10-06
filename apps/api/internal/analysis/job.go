package analysis

import (
	"log"

	"github.com/codeatlas/api/analyzer"
)

func (s *Service) runJob(run *AnalysisRun, repoPath string) {
	go func() {
		s.store.updateStatus(run.ID, "running", "")

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

		s.store.updateStatus(run.ID, "completed", "")
	}()
}

func (s *Service) persistResult(runID uint, result *analyzer.Result) error {
	// TODO: persist result.Files, result.Symbols, result.Dependencies, result.Endpoints
	return nil
}
