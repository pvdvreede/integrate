package main

import (
	"github.com/op/go-logging"
	"github.com/pvdvreede/integrate"
	"log"
	"net/http"
)

type LoggerHandler struct{}

func (l *LoggerHandler) ShouldAction(m *integrate.Message, logger integrate.Logger) bool {
	logger.Debug("Received should action for message %v", m)
	return true
}

func (l *LoggerHandler) Action(m *integrate.Message, logger integrate.Logger) (*integrate.Message, error) {
	logger.Debug("Received Action for message %v", m)
	newMessage := m.Copy()
	newMessage.AddContext("logged", true)
	return newMessage, nil
}

func main() {
	logger := logging.MustGetLogger("example")
	storage := &integrate.ProbeStorage{}
	server := integrate.NewServer("http_test", logger, storage)

	server.AddHandlers(&LoggerHandler{})
	comms, err := server.Run()

	if err != nil {
		panic(err)
	}

	http.HandleFunc("/input/test1", func(w http.ResponseWriter, r *http.Request) {
		c := integrate.Context{
			"adapter":   "http",
			"http_path": r.RequestURI,
		}
		m := integrate.NewMessage(r.Body, c)
		comms <- m
		w.Header().Add("X-TransactionID", m.TransactionId)
		w.Write([]byte("received"))
	})

	logger.Notice("Starting HTTP server on port 8085...")
	log.Fatal(http.ListenAndServe(":8085", nil))
}
