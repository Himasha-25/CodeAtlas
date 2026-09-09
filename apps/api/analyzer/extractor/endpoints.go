package extractor

import "github.com/codeatlas/api/analyzer/model"

// ExtractEndpoints detects HTTP API endpoint definitions from symbols.
// v1: pattern-match on symbol names (e.g. gin route registrations).
func ExtractEndpoints(symbols []model.Symbol) []model.Endpoint {
	// TODO: implement pattern matching for common HTTP frameworks
	return nil
}
