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

// Used to probe which methods have been called. Main use is for testing.
type ProbeStorage struct {
	HasActionedCallCount   int
	StartProcessCallCount  int
	FinishProcessCallCount int
	StoreCallCount         int
	SetHasActioned         bool
	Messages               []*Message
}

// Will return the value of the property SetHasActioned. This allows us to test
// for different scenarios.
func (s *ProbeStorage) HasActioned(m *Message, h Handler) bool {
	s.HasActionedCallCount += 1
	return s.SetHasActioned
}

// a NoOp that always returns NO error.
func (s *ProbeStorage) StartProcess(m *Message) error {
	s.StartProcessCallCount += 1
	return nil
}

// a NoOp that always returns NO error.
func (s *ProbeStorage) FinishProcess(m *Message) error {
	s.FinishProcessCallCount += 1
	return nil
}

// a NoOp that always returns NO error.
func (s *ProbeStorage) Store(m *Message, h Handler, event string) error {
	s.StoreCallCount += 1
	s.Messages = append(s.Messages, m)
	return nil
}
