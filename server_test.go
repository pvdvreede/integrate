package integrate

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestServer(t *testing.T) {

	Convey("NewServer", t, func() {
		logger := &NoOpLogger{}
		storage := &ProbeStorage{}
		s := NewServer("test", logger, storage)

		Convey("Returns a pointer to a server instance", func() {
			So(s, ShouldHaveSameTypeAs, &Server{})
		})

		Convey("Sets the name", func() {
			So(s.name, ShouldEqual, "test")
		})

		Convey("Sets the logger", func() {
			So(s.logger, ShouldEqual, logger)
		})
		Convey("Sets the storage", func() {
			So(s.storage, ShouldEqual, storage)
		})
	})

	Convey("Run", t, func() {
		logger := &NoOpLogger{}
		storage := &ProbeStorage{}
		s := NewServer("test", logger, storage)
		comms, err := s.Run()
		s.Stop()

		Convey("Returns a chan of Message pointers", func() {
			So(comms, ShouldHaveSameTypeAs, make(chan *Message))
		})

		Convey("Should not return an error", func() {
			So(err, ShouldBeNil)
		})
	})

	Convey("processMessage", t, func() {
		logger := &NoOpLogger{}
		storage := &ProbeStorage{}
		s := NewServer("test", logger, storage)
		ph1 := &ProbeHandler{SetShouldAction: true}
		ph2 := &ProbeHandler{SetShouldAction: false}
		ph3 := &ProbeHandler{SetShouldAction: true}
		message := &Message{}
		s.AddHandlers(ph1, ph2, ph3)
		s.processMessage(message)

		Convey("Only runs handlers returning true for ShouldAction", func() {
			So(ph1.ActionCallCount, ShouldEqual, 1)
			So(ph2.ActionCallCount, ShouldEqual, 0)
			So(ph3.ActionCallCount, ShouldEqual, 1)
		})

		Convey("Calls ShouldAction for every handler", func() {
			So(ph1.ShouldActionCallCount, ShouldEqual, 1)
			So(ph2.ShouldActionCallCount, ShouldEqual, 1)
			So(ph3.ShouldActionCallCount, ShouldEqual, 1)
		})

		Convey("Calls HasActioned on storage for every handler", func() {
			So(storage.HasActionedCallCount, ShouldEqual, 3)
		})

		Convey("Call StartProcess on storage", func() {
			So(storage.StartProcessCallCount, ShouldEqual, 1)
		})

		Convey("Call FinishProcess on storage", func() {
			So(storage.FinishProcessCallCount, ShouldEqual, 1)
		})

		Convey("Call Store twice on every actioned handler", func() {
			So(storage.StoreCallCount, ShouldEqual, 4)
		})

		Convey("Store a new message on every post action Store call", func() {
			So(storage.Messages[0].Id, ShouldEqual, message.Id)
			So(storage.Messages[1].Id, ShouldNotEqual, message.Id)
			So(storage.Messages[2].Id, ShouldEqual, storage.Messages[1].Id)
			So(storage.Messages[3].Id, ShouldNotEqual, storage.Messages[2].Id)
		})

		Convey("Keeps the same transaction Id for each message", func() {
			transactionId := message.TransactionId
			So(storage.Messages[0].TransactionId, ShouldEqual, transactionId)
			So(storage.Messages[1].TransactionId, ShouldEqual, transactionId)
			So(storage.Messages[2].TransactionId, ShouldEqual, transactionId)
			So(storage.Messages[3].TransactionId, ShouldEqual, transactionId)
		})
	})
}
