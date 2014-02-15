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
	Id      string
	Context Context
	Data    io.Reader
}

// Add some context to the message
func (m *Message) AddContext(key string, value interface{}) {
	m.Context[key] = value
}

// Implement stringer interface to render a message as a string for logging
// and storing purposes.
func (m *Message) String() string {
	return m.Id
}

// Create a new base message with data and an initial context.
func NewMessage(data io.Reader, context Context) *Message {
	message := Message{
		Data:    data,
		Context: context,
		Id:      generateMessageId(),
	}
	return &message
}

func generateMessageId() string {
	guid := uuid.NewRandom()
	return string(guid)
}
