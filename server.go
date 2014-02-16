package integrate

// The actual integrate server that provides a developer API for integrating.
type Server struct {
	name      string
	logger    Logger
	storage   Storage
	handlers  []Handler
	commsChan chan *Message
	killChan  chan bool
}

// Start the integrate server up. This will spawn a new go routine that will
// handle the message processing, and immeadiately return with a channel to
// send messages on.
func (s *Server) Run() (chan *Message, error) {
	s.logger.Notice("Starting %v server...", s.name)
	s.commsChan = make(chan *Message)
	s.killChan = make(chan bool)
	go s.listenAndServe()
	return s.commsChan, nil
}

// Stop the server from running. This will send the kill command to the
// core go routine and close the channels the server uses. So the comms channel
// will through errors if it is used after calling this command.
func (s *Server) Stop() {
	s.logger.Notice("Shutting down %v server...", s.name)
	s.killChan <- true
	s.logger.Notice("Server %v shutdown complete.", s.name)
}

// Add handlers that will be run for each message.
// The handlers are run in order, serially.
func (s *Server) AddHandlers(handlers ...Handler) {
	for _, h := range handlers {
		s.handlers = append(s.handlers, h)
	}
}

func (s *Server) listenAndServe() {
	for {
		select {
		case m := <-s.commsChan:
			s.logger.Debug("Recieved message %v to be processed...", m)
			go s.processMessage(m)
		case <-s.killChan:
			s.logger.Debug("Kill command received, so exiting loop.")
			break
		}
	}

	close(s.commsChan)
	close(s.killChan)
}

func (s *Server) processMessage(incoming *Message) {
	workingMessage := incoming
	s.storage.StartProcess(workingMessage)
	for _, h := range s.handlers {
		if !s.storage.HasActioned(incoming, h) && h.ShouldAction(incoming, s.logger) {
			s.storage.Store(workingMessage, h, "pre-action")
			currentMessage, _ := h.Action(workingMessage, s.logger)
			workingMessage = currentMessage
			s.storage.Store(workingMessage, h, "post-action")
		}
	}
	s.storage.FinishProcess(workingMessage)
}

// Create a new server instance with user's choice for logging and storage.
// A name is used in logging and process storage to differentiate from other
// integrate servers running in the same process. This allows you to have
// a secure multi-tenant system of different integrate servers running together
// but safely segregate them.
// The logger and storage can be the same across multiple integrate servers
// (assuming they are thread safe) or you can use a different one for each
// for more segregation.
func NewServer(name string, log Logger, store Storage) *Server {
	s := Server{
		name:    name,
		logger:  log,
		storage: store,
	}

	return &s
}
