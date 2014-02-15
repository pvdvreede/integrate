package integrate

import (
	. "github.com/smartystreets/goconvey/convey"
	"io"
	"testing"
)

func TestMessage(t *testing.T) {
	context := &Context{}
	data, _ := io.Pipe()
	message := NewMessage(data, context)

	Convey("NewMessage", t, func() {

		Convey("Set a unique id", func() {
			So(message.Id, ShouldNotBeNil)
			So(len(message.Id), ShouldEqual, 16)
		})

		Convey("Set the data attribute", func() {
			So(message.Data, ShouldEqual, data)
		})

		Convey("Set the context object", func() {
			So(message.Context, ShouldEqual, context)
		})
	})

	Convey("String", t, func() {
		Convey("Returns the Id", func() {
			So(message.String(), ShouldEqual, message.Id)
		})
	})
}
