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

func Test_Iterator(t *testing.T) {
	a := NewSet()

	a.Add("Z")
	a.Add("Y")
	a.Add("X")
	a.Add("W")

	b := NewSet()
	for val := range a.Iterator().C {
		b.Add(val)
	}

	if !a.Equal(b) {
		t.Error("The sets are not equal after iterating (Iterator) through the first set")
	}
}
