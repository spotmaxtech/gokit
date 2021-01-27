package gokit

import (
	"fmt"
	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestLogrusBufferHook(t *testing.T) {
	Convey("logrus buffer hook", t, func() {
		hook, buffer := NewLogrusBufferHook([]logrus.Level{logrus.WarnLevel, logrus.InfoLevel}, 5)

		logrus.AddHook(hook)

		logrus.WithFields(logrus.Fields{
			"name": "spotmax",
			"age":  15,
		}).Warn("Hello world! 1")

		logrus.Info("Hello world! 2")
		logrus.Debug("Hello world! 3")
		logrus.Info("Hello world! 4")
		logrus.Info("Hello world! 5")
		logrus.Info("Hello world! 6")
		logrus.Info("Hello world! 7")
		logrus.WithFields(logrus.Fields{
			"name": "spotmax",
			"age":  16,
		}).Warn("Hello world! 8")

		logs := buffer.Read()
		for _, l := range logs {
			fmt.Println(l)
		}
	})
}
