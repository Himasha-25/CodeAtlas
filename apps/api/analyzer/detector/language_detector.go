package detector

import "path/filepath"

var extToLang = map[string]string{
	".go":  "go",
	".ts":  "typescript",
	".tsx": "typescript",
	".js":  "javascript",
	".jsx": "javascript",
	".py":  "python",
}

// Detect returns the programming language for a file path based on extension.
func Detect(path string) string {
	lang, ok := extToLang[filepath.Ext(path)]
	if !ok {
		return "unknown"
	}
	return lang
}
