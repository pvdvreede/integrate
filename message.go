package integrate

import (
	"code.google.com/p/go-uuid/uuid"
	"io"
)

// Storage of key value pairs that are used to provide context about the
// message.
type Context map[string]interface{}

// Core data structure for each message being passed through the system.
type Message struct {
	Id            string
	ParentId      string
	TransactionId string
	Context       Context
	Data          io.Reader
}

// Add some context to the message
func (m *Message) AddContext(key string, value interface{}) {
	m.Context[key] = value
}

// Creates a new message with the same data and context as the current message.
// It also sets the parentId as the current message and the same transaction id.
func (m *Message) Copy() *Message {
	newMessage := NewMessage(m.Data, m.Context)
	newMessage.ParentId = m.Id
	newMessage.TransactionId = m.TransactionId
	return newMessage
}

// Implement stringer interface to render a message as a string for logging
// and storing purposes.
func (m *Message) String() string {
	return m.Id
}

// Create a new base message with data and an initial context.
func NewMessage(data io.Reader, context Context) *Message {
	message := Message{
		Data:          data,
		Context:       context,
		Id:            generateMessageId(),
		TransactionId: generateMessageId(),
	}
	return &message
}

func generateMessageId() string {
	return uuid.New()
}
