package integrate

import (
	. "github.com/smartystreets/goconvey/convey"
	"io"
	"regexp"
	"testing"
)

func TestMessage(t *testing.T) {
	context := Context{}
	data, _ := io.Pipe()
	message := NewMessage(data, context)

	Convey("NewMessage", t, func() {

		Convey("Set a unique id", func() {
			So(message.Id, ShouldNotBeNil)
			So(len(message.Id), ShouldEqual, 36)
		})

		Convey("Set a unique transactionId", func() {
			So(message.TransactionId, ShouldNotBeNil)
			So(len(message.TransactionId), ShouldEqual, 36)
		})

		Convey("Set the data attribute", func() {
			So(message.Data, ShouldEqual, data)
		})

		Convey("Set the context object", func() {
			So(message.Context, ShouldEqual, context)
		})

		Convey("Id's should be v4 GUIDs", func() {
			reg := regexp.MustCompile(`^[A-Za-z0-9]{8}-[A-Za-z0-9]{4}-[A-Za-z0-9]{4}-[A-Za-z0-9]{4}-[A-Za-z0-9]{12}$`)
			So(reg.MatchString(message.Id), ShouldBeTrue)
			So(reg.MatchString(message.TransactionId), ShouldBeTrue)
		})
	})

	Convey("String", t, func() {
		Convey("Returns the Id", func() {
			So(message.String(), ShouldEqual, message.Id)
		})
	})

	Convey("AddContext", t, func() {
		message.AddContext("key", "value")

		Convey("Adds to the context property", func() {
			So(message.Context["key"], ShouldEqual, "value")
		})
	})

	Convey("Copy", t, func() {
		newMessage := message.Copy()

		Convey("Returns a new message", func() {
			So(newMessage, ShouldNotEqual, message)
		})

		Convey("Has a new Id set", func() {
			So(newMessage.Id, ShouldNotBeNil)
			So(len(newMessage.Id), ShouldEqual, 36)
		})

		Convey("Has the correct ParentId set", func() {
			So(newMessage.ParentId, ShouldEqual, message.Id)
		})

		Convey("Has the correct TransactionId set", func() {
			So(newMessage.TransactionId, ShouldEqual, message.TransactionId)
		})

		Convey("Copies the data", func() {
			So(newMessage.Data, ShouldEqual, message.Data)
		})

		Convey("Copies the context", func() {
			So(newMessage.Context, ShouldEqual, message.Context)
		})
	})
}
