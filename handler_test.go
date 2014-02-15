package integrate

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestHandler(t *testing.T) {
	Convey("NoOpHandler", t, func() {
		logger := &NoOpLogger{}
		noOpHandler := &NoOpHandler{}

		Convey("Should return true for ShouldAction", func() {
			So(noOpHandler.ShouldAction(&Message{}, logger), ShouldBeTrue)
		})

		Convey("Should return new message and no error", func() {
			oldMsg := &Message{}
			newMsg, err := noOpHandler.Action(oldMsg, logger)
			So(oldMsg, ShouldNotEqual, newMsg)
			So(err, ShouldBeNil)
		})
	})
}
