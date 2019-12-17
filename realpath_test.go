package gokit

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRealpath(t *testing.T) {
	Convey("test realpath", t, func() {
		t.Log(Realpath("./config.json"))
		t.Log(Realpath("./../config.json"))
	})
}
