package assistant

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Service struct {
	store     *store
	retriever Retriever
	llm       LLMClient
}

func NewService(db *gorm.DB, groqAPIKey string) *Service {
	return &Service{
		store:     newStore(db),
		retriever: newRetriever(db),
		llm:       NewGroqClient(groqAPIKey),
	}
}

type AskRequest struct {
	ConversationID uint   `json:"conversationId"`
	RunID          uint   `json:"runId"`
	Question       string `json:"question" binding:"required"`
}

type AskResponse struct {
	Answer  string    `json:"answer"`
	Sources []Snippet `json:"sources"`
}

func (s *Service) Ask(projectID uint, req AskRequest) (*AskResponse, error) {
	snippets, err := s.retriever.Retrieve(req.RunID, req.Question, 5)
	if err != nil {
		return nil, err
	}
	prompt := buildPrompt(req.Question, snippets)
	answer, err := s.llm.Complete(prompt)
	if err != nil {
		return nil, err
	}
	sourcesJSON, _ := json.Marshal(snippets)
	_ = s.store.saveMessage(&Message{
		ConversationID: req.ConversationID,
		Role:           "user",
		Content:        req.Question,
	})
	_ = s.store.saveMessage(&Message{
		ConversationID: req.ConversationID,
		Role:           "assistant",
		Content:        answer,
		Sources:        string(sourcesJSON),
	})
	return &AskResponse{Answer: answer, Sources: snippets}, nil
}
