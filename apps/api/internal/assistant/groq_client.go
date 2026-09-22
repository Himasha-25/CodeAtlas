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

const groqEndpoint = "https://api.groq.com/openai/v1/chat/completions"
const groqModel = "llama-3.3-70b-versatile"

type groqClient struct {
	apiKey string
}

func NewGroqClient(apiKey string) LLMClient {
	return &groqClient{apiKey: apiKey}
}

func (c *groqClient) Complete(prompt string) (string, error) {
	body, _ := json.Marshal(map[string]interface{}{
		"model": groqModel,
		"messages": []map[string]string{
			{"role": "user", "content": prompt},
		},
	})

	req, err := http.NewRequest(http.MethodPost, groqEndpoint, bytes.NewReader(body))
	if err != nil {
		return "", fmt.Errorf("groq: build request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("groq: request: %w", err)
	}
	defer resp.Body.Close()

	raw, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("groq: non-200 status %d: %s", resp.StatusCode, raw)
	}

	var result struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}
	if err := json.Unmarshal(raw, &result); err != nil {
		return "", fmt.Errorf("groq: parse response: %w", err)
	}
	if len(result.Choices) == 0 {
		return "", fmt.Errorf("groq: empty choices in response")
	}
	return result.Choices[0].Message.Content, nil
}
