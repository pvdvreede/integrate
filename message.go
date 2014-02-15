package integrate

import (
	"code.google.com/p/go-uuid/uuid"
	"io"
)

// Key value structure for storing a value which represents a single piece of
// the message's context.
type Metadata struct {
	key   string
	value interface{}
}

// Storage of key value pairs that are used to provide context about the
// message.
type Context []Metadata

// Core data structure for each message being passed through the system.
type Message struct {
	Id      string
	Context *Context
	Data    io.Reader
}

// Implement stringer interface to render a message as a string for logging
// and storing purposes.
func (m *Message) String() string {
	return m.Id
}

// Create a new base message with data and an initial context.
func NewMessage(data io.Reader, context *Context) *Message {
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
