package parser

import "github.com/codeatlas/api/analyzer/model"

type jsParser struct{}

func (p *jsParser) Parse(path string) ([]model.Symbol, error) {
	// TODO: implement with tree-sitter javascript grammar
	return nil, nil
}
