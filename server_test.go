package integrate

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestServer(t *testing.T) {
	logger := &NoOpLogger{}
	storage := &NoOpStorage{}
	s := NewServer("test", logger, storage)

	Convey("NewServer", t, func() {

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
		comms, err := s.Run()

		Convey("Returns a chan of Message pointers", func() {
			So(comms, ShouldHaveSameTypeAs, make(chan *Message))
		})

		Convey("Should not return an error", func() {
			So(err, ShouldBeNil)
		})
	})
}
