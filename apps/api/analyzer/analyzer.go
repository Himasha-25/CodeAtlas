package analyzer

import (
	"github.com/codeatlas/api/analyzer/dependency"
	"github.com/codeatlas/api/analyzer/detector"
	"github.com/codeatlas/api/analyzer/extractor"
	"github.com/codeatlas/api/analyzer/model"
	"github.com/codeatlas/api/analyzer/scanner"
)

// Result is the full output of an analysis run.
type Result struct {
	Files        []model.FileInfo
	Symbols      []model.Symbol
	Dependencies []model.Dependency
	Endpoints    []model.Endpoint
}

// Analyze runs the full analysis pipeline on a repository at repoPath.
func Analyze(repoPath string) (*Result, error) {
	paths, err := scanner.Scan(repoPath)
	if err != nil {
		return nil, err
	}

	files := make([]model.FileInfo, 0, len(paths))
	for _, p := range paths {
		files = append(files, model.FileInfo{
			Path:     p,
			Language: detector.Detect(p),
		})
	}

	symbols, err := extractor.ExtractSymbols(files)
	if err != nil {
		return nil, err
	}

	deps := dependency.Build(files, symbols)
	endpoints := extractor.ExtractEndpoints(symbols)

	return &Result{
		Files:        files,
		Symbols:      symbols,
		Dependencies: deps,
		Endpoints:    endpoints,
	}, nil
}
