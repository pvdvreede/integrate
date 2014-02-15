package integrate

// This is the interface that must be satisfied for all Actions that the user
// wants to implement.
type Handler interface {

	// This method will return true if the actual Action should be peformed for
	// the passed in Message or false if it should be skipped. This can be used
	// to cut down logging and storing as there will be no storage events if
	// this returns false.
	ShouldAction(m *Message, logger Logger) bool

	// This is the actual worker function that performs the work for each message.
	// It returns a new message that will be passed to the next Handler in the
	// list. If error is not nil, it will halt processing the rest of the
	// handlers and put the Message in an error'd state for review.
	Action(m *Message, logger Logger) (*Message, error)
}

// This is a handler that will always run, but does absolutely nothing. Can
// be used for testing or as a placeholder.
type NoOpHandler struct{}

// Will always return true
func (n *NoOpHandler) ShouldAction(m *Message, logger Logger) bool {
	return true
}

// Will always return a new Message pointer with the old data and context and
// always return nil for the error; so it will never fail.
func (n *NoOpHandler) Action(m *Message, logger Logger) (*Message, error) {
	return NewMessage(m.Data, m.Context), nil
}
