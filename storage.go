package integrate

// This is the interface used for implementing a storage mechanism to store
// messages as they are running through the system. This data can then be used
// for trouble shooting issues or tracking messages as well as restarting
// error'd messages.
type Storage interface {

	// This is called by the core to make sure a message doesn't get
	// run through a handler twice after a re-process. It should
	// return true if this message has already been run through this handler.
	HasActioned(m *Message, h Handler) bool

	// Called when a new process is starting. Should return error if there is a
	// problem storing the data.
	StartProcess(m *Message) error

	// Called when a process has run all the handlers and finished successfull.
	// Should only return an error if there is an issue persisting the data.
	FinishProcess(m *Message) error

	// The main function that should be used to store details into a persistent
	// storage of the implementers choice. The event string will highlight where
	// the message is in the context of the handler.
	// This will return an error if there is an issue storing the data.
	Store(m *Message, h Handler, event string) error
}

// A storage that does nothing. It should be used for testing or demoing, when
// persistent storage or process/message data is not needed. You can tell the
// storage if you want to return true or not for HasActioned with the property
// SetHasActioned.
type NoOpStorage struct {
	SetHasActioned bool
}

// Will return the value of the property SetHasActioned. This allows us to test
// for different scenarios.
func (s *NoOpStorage) HasActioned(m *Message, h Handler) bool {
	return s.SetHasActioned
}

// a NoOp that always returns NO error.
func (s *NoOpStorage) StartProcess(m *Message) error {
	return nil
}

// a NoOp that always returns NO error.
func (s *NoOpStorage) FinishProcess(m *Message) error {
	return nil
}

// a NoOp that always returns NO error.
func (s *NoOpStorage) Store(m *Message, h Handler, event string) error {
	return nil
}
