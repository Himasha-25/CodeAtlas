package parser

import "github.com/codeatlas/api/analyzer/model"

// Parser extracts symbols from a source file.
type Parser interface {
	Parse(path string) ([]model.Symbol, error)
}

// For returns the parser for a given language, or nil if unsupported.
func For(language string) Parser {
	switch language {
	case "go":
		return &goParser{}
	case "typescript":
		return &tsParser{}
	case "javascript":
		return &jsParser{}
	default:
		return nil
	}
}
