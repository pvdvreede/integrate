package integrate

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestLogger(t *testing.T) {
	Convey("MemoryLogger", t, func() {
		mlogger := &MemoryLogger{}

		Convey("Debug logs to Debugs slice", func() {
			mlogger.Debug("This is log text for %v", "Debug")
			So(mlogger.Debugs[0], ShouldEqual, "This is log text for [Debug]")
		})

		Convey("Notice logs to Notices slice", func() {
			mlogger.Notice("This is log text for %v", "Notice")
			So(mlogger.Notices[0], ShouldEqual, "This is log text for [Notice]")
		})

		Convey("Warning logs to Warnings slice", func() {
			mlogger.Warning("This is log text for %v", "Warning")
			So(mlogger.Warnings[0], ShouldEqual, "This is log text for [Warning]")
		})

		Convey("Error logs to Errors slice", func() {
			mlogger.Error("This is log text for %v", "Error")
			So(mlogger.Errors[0], ShouldEqual, "This is log text for [Error]")
		})

		Convey("Critical logs to Criticals slice", func() {
			mlogger.Critical("This is log text for %v", "Critical")
			So(mlogger.Criticals[0], ShouldEqual, "This is log text for [Critical]")
		})
	})
}
