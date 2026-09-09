package assistant

import "gorm.io/gorm"

// Retriever finds relevant code snippets for a query.
type Retriever interface {
	Retrieve(runID uint, query string, topK int) ([]Snippet, error)
}

type Snippet struct {
	FilePath string
	Symbol   string
	Content  string
}

type postgresRetriever struct{ db *gorm.DB }

func newRetriever(db *gorm.DB) Retriever { return &postgresRetriever{db} }

func (r *postgresRetriever) Retrieve(runID uint, query string, topK int) ([]Snippet, error) {
	// v1: simple ILIKE search over symbols joined with files
	rows, err := r.db.Raw(`
		SELECT f.path, s.name, s.name AS content
		FROM symbols s
		JOIN files f ON f.id = s.file_id
		WHERE s.run_id = ? AND s.name ILIKE ?
		LIMIT ?`, runID, "%"+query+"%", topK).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var snippets []Snippet
	for rows.Next() {
		var sn Snippet
		rows.Scan(&sn.FilePath, &sn.Symbol, &sn.Content)
		snippets = append(snippets, sn)
	}
	return snippets, nil
}
