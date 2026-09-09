package assistant

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// LLMClient sends a prompt and returns a response.
type LLMClient interface {
	Complete(prompt string) (string, error)
}

type ollamaClient struct {
	baseURL string
	model   string
}

func newOllamaClient(baseURL, model string) LLMClient {
	return &ollamaClient{baseURL: baseURL, model: model}
}

func (c *ollamaClient) Complete(prompt string) (string, error) {
	body, _ := json.Marshal(map[string]interface{}{
		"model":  c.model,
		"prompt": prompt,
		"stream": false,
	})
	resp, err := http.Post(c.baseURL+"/api/generate", "application/json", bytes.NewReader(body))
	if err != nil {
		return "", fmt.Errorf("ollama request: %w", err)
	}
	defer resp.Body.Close()
	raw, _ := io.ReadAll(resp.Body)
	var result struct {
		Response string `json:"response"`
	}
	if err := json.Unmarshal(raw, &result); err != nil {
		return "", fmt.Errorf("ollama parse: %w", err)
	}
	return result.Response, nil
}
