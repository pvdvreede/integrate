package integrate

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestServer(t *testing.T) {

	Convey("NewServer", t, func() {
		logger := &NoOpLogger{}
		storage := &NoOpStorage{}
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
		storage := &NoOpStorage{}
		s := NewServer("test", logger, storage)
		comms, err := s.Run()

		Convey("Returns a chan of Message pointers", func() {
			So(comms, ShouldHaveSameTypeAs, make(chan *Message))
		})

		Convey("Should not return an error", func() {
			So(err, ShouldBeNil)
		})
	})

	Convey("processMessage", t, func() {
		logger := &NoOpLogger{}
		storage := &NoOpStorage{}
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

		Convey("Calls HasActioned on storage for every handler", nil)
	})
}
