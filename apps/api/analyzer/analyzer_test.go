package analyzer_test

import (
	"testing"

	"github.com/codeatlas/api/analyzer"
)

func TestAnalyze_SampleGoProject(t *testing.T) {
	result, err := analyzer.Analyze("../testdata/sample-go-project")
	if err != nil {
		t.Fatalf("Analyze failed: %v", err)
	}
	if len(result.Files) == 0 {
		t.Error("expected files, got none")
	}
}

func TestAnalyze_SampleTSProject(t *testing.T) {
	result, err := analyzer.Analyze("../testdata/sample-ts-project")
	if err != nil {
		t.Fatalf("Analyze failed: %v", err)
	}
	_ = result
}
