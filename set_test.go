package gokit

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSet(t *testing.T) {
	Convey("test common use", t, func() {
		set := NewSet()
		set.Add("a")
		set.Add("b")
		So(set.Cardinality(), ShouldEqual, 2)
		So(set.Contains("a"), ShouldBeTrue)
	})
}
