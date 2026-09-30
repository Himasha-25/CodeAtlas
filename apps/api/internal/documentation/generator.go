package docs

import "fmt"

func generate(projectID, runID uint) string {
	return fmt.Sprintf("# Project %d Documentation\n\nGenerated from analysis run %d.\n", projectID, runID)
}
