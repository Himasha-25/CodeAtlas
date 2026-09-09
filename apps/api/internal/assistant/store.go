package assistant

import (
	"time"

	"gorm.io/gorm"
)

type Conversation struct {
	ID        uint      `gorm:"primaryKey"`
	ProjectID uint      `gorm:"not null;index"`
	CreatedAt time.Time
}

type Message struct {
	ID             uint      `gorm:"primaryKey"`
	ConversationID uint      `gorm:"not null;index"`
	Role           string    `gorm:"not null"` // user|assistant
	Content        string    `gorm:"type:text;not null"`
	Sources        string    `gorm:"type:text"` // JSON array of file/symbol refs
	CreatedAt      time.Time
}

type store struct{ db *gorm.DB }

func newStore(db *gorm.DB) *store { return &store{db} }

func (s *store) createConversation(c *Conversation) error { return s.db.Create(c).Error }

func (s *store) saveMessage(m *Message) error { return s.db.Create(m).Error }

func (s *store) listMessages(conversationID uint) ([]Message, error) {
	var msgs []Message
	return msgs, s.db.Where("conversation_id = ?", conversationID).Order("created_at asc").Find(&msgs).Error
}
