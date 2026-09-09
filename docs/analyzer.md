# Analyzer

Go package at `apps/api/analyzer/`. Zero dependencies on `internal/*` — testable with `go test ./analyzer/...` and no DB or server running.

## Pipeline

```
scanner.Scan()       → file list (excludes node_modules, .git, vendor)
detector.Detect()    → language per file
parser.Parse()       → tree-sitter AST per file
extractor.Extract()  → symbols (functions, classes, interfaces, imports/exports)
dependency.Build()   → file-to-file and symbol-to-symbol edges
→ model.Result
```

## Adding a Language

1. Add a parser implementation in `analyzer/parser/<lang>.go` implementing the `Parser` interface.
2. Wire it into `detector/language_detector.go`.
3. Add test fixtures in `testdata/sample-<lang>-project/`.
