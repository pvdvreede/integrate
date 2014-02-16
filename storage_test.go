package integrate

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestStorage(t *testing.T) {
	Convey("ProbeStorage", t, func() {
		probeStorage := &ProbeStorage{}
		noOpHandler := &NoOpHandler{}
		message := &Message{}

		Convey("Should return SetHasActioned", func() {
			probeStorage.SetHasActioned = true
			So(probeStorage.HasActioned(message, noOpHandler), ShouldBeTrue)

			probeStorage.SetHasActioned = false
			So(probeStorage.HasActioned(message, noOpHandler), ShouldBeFalse)
		})

		Convey("Should return nil for StartProcess", func() {
			So(probeStorage.StartProcess(message), ShouldBeNil)
		})

		Convey("Should return nil for FinishProcess", func() {
			So(probeStorage.FinishProcess(message), ShouldBeNil)
		})
	})
}
