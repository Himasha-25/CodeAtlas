package documentation

import "fmt"

// generate builds markdown documentation from symbol/file data.
// In a real implementation this would query codeexplorer store and format results.
func generate(projectID, runID uint) string {
	return fmt.Sprintf("# Project %d Documentation\n\nGenerated from analysis run %d.\n", projectID, runID)
}
