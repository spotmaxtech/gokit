package gokit

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestSimpleTimer_Timeout(t *testing.T) {
	Convey("test timeout", t, func() {
		timer := NewSimpleTimer(time.Millisecond * 100)
		So(timer.Timeout(), ShouldBeFalse)
		time.Sleep(time.Millisecond * 150)
		So(timer.Timeout(), ShouldBeTrue)
		time.Sleep(time.Millisecond * 150)
		So(timer.Timeout(), ShouldBeTrue)
	})
}

func TestSimpleTimer_Reset(t *testing.T) {
	Convey("test timeout", t, func() {
		timer := NewSimpleTimer(time.Millisecond * 100)
		So(timer.Timeout(), ShouldBeFalse)
		time.Sleep(time.Millisecond * 150)
		So(timer.Timeout(), ShouldBeTrue)
		timer.Reset()
		So(timer.Timeout(), ShouldBeFalse)
		time.Sleep(time.Millisecond * 150)
		So(timer.Timeout(), ShouldBeTrue)
	})
}

func TestSimpleTimer_Checkpoint(t *testing.T) {
	Convey("test timeout", t, func() {
		timer := NewSimpleTimer(time.Millisecond * 100)
		So(timer.Checkpoint(), ShouldBeFalse)
		time.Sleep(time.Millisecond * 150)
		So(timer.Checkpoint(), ShouldBeTrue)
		So(timer.Checkpoint(), ShouldBeFalse)
		So(timer.Checkpoint(), ShouldBeFalse)
		time.Sleep(time.Millisecond * 150)
		So(timer.Checkpoint(), ShouldBeTrue)
	})
}
