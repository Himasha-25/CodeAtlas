package impact

import "gorm.io/gorm"

type Edge struct {
	FromID   uint
	ToID     uint
	EdgeType string
}

// traverse returns all nodes reachable from fromID in the dependency graph for a given run.
func traverse(db *gorm.DB, runID, fromID uint) ([]uint, error) {
	visited := map[uint]bool{fromID: true}
	queue := []uint{fromID}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		var edges []Edge
		db.Raw("SELECT to_id FROM dependencies WHERE run_id = ? AND from_id = ?", runID, current).Scan(&edges)
		for _, e := range edges {
			if !visited[e.ToID] {
				visited[e.ToID] = true
				queue = append(queue, e.ToID)
			}
		}
	}
	result := make([]uint, 0, len(visited))
	for id := range visited {
		result = append(result, id)
	}
	return result, nil
}
