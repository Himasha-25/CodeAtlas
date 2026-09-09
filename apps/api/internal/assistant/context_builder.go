package assistant

import (
	"fmt"
	"strings"
)

func buildPrompt(question string, snippets []Snippet) string {
	var sb strings.Builder
	sb.WriteString("You are a code assistant. Answer based only on the provided code context.\n\n")
	sb.WriteString("## Code Context\n\n")
	for _, s := range snippets {
		sb.WriteString(fmt.Sprintf("// %s — %s\n%s\n\n", s.FilePath, s.Symbol, s.Content))
	}
	sb.WriteString("## Question\n\n")
	sb.WriteString(question)
	return sb.String()
}
