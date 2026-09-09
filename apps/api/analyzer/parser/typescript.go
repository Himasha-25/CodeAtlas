package parser

import "github.com/codeatlas/api/analyzer/model"

type tsParser struct{}

func (p *tsParser) Parse(path string) ([]model.Symbol, error) {
	// TODO: implement with tree-sitter typescript grammar
	return nil, nil
}
