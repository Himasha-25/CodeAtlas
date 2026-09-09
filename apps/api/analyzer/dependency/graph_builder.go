package dependency

import "github.com/codeatlas/api/analyzer/model"

// Build constructs file-to-file dependency edges from import/require statements in symbols.
func Build(files []model.FileInfo, symbols []model.Symbol) []model.Dependency {
	// TODO: implement import resolution per language
	return nil
}
