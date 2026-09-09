package parser

import "github.com/codeatlas/api/analyzer/model"

type goParser struct{}

func (p *goParser) Parse(path string) ([]model.Symbol, error) {
	// TODO: implement with tree-sitter go grammar
	return nil, nil
}
