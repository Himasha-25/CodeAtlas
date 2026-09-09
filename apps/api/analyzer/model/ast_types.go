package model

// FileInfo represents a source file discovered in the repo.
type FileInfo struct {
	Path      string
	Language  string
	LineCount int
}

// Symbol represents a named code entity extracted from a file.
type Symbol struct {
	Name     string
	Kind     string // function|class|interface|variable
	File     string
	Line     int
	Exported bool
}

// Dependency is a directed edge between two files or symbols.
type Dependency struct {
	FromFile string
	ToFile   string
	Kind     string // import|call|inherit
}

// Endpoint is a detected HTTP API endpoint.
type Endpoint struct {
	Method string
	Path   string
	File   string
	Line   int
}
