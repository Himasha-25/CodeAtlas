package extractor

import (
	"github.com/codeatlas/api/analyzer/model"
	"github.com/codeatlas/api/analyzer/parser"
)

// ExtractSymbols runs the appropriate parser for each file and returns all symbols.
func ExtractSymbols(files []model.FileInfo) ([]model.Symbol, error) {
	var all []model.Symbol
	for _, f := range files {
		p := parser.For(f.Language)
		if p == nil {
			continue
		}
		symbols, err := p.Parse(f.Path)
		if err != nil {
			return nil, err
		}
		all = append(all, symbols...)
	}
	return all, nil
}
