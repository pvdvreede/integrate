package integrate

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestStorage(t *testing.T) {
	Convey("NoOpStorage", t, func() {
		noOpStorage := &NoOpStorage{}
		noOpHandler := &NoOpHandler{}
		message := &Message{}

		Convey("Should return SetHasActioned", func() {
			noOpStorage.SetHasActioned = true
			So(noOpStorage.HasActioned(message, noOpHandler), ShouldBeTrue)

			noOpStorage.SetHasActioned = false
			So(noOpStorage.HasActioned(message, noOpHandler), ShouldBeFalse)
		})

		Convey("Should return nil for StartProcess", func() {
			So(noOpStorage.StartProcess(message), ShouldBeNil)
		})

		Convey("Should return nil for FinishProcess", func() {
			So(noOpStorage.FinishProcess(message), ShouldBeNil)
		})
	})
}
