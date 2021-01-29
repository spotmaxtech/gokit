package gokit

import (
	"fmt"
	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"testing"
	"time"
)

func TestLogrusBufferHook(t *testing.T) {
	Convey("logrus buffer hook", t, func() {
		hook, buffer := NewLogrusBufferHook([]logrus.Level{logrus.WarnLevel, logrus.InfoLevel}, 5)

		logrus.AddHook(hook)

		logrus.Warn("Hello world! 1")
		logrus.Info("Hello world! 2")
		logrus.Debug("Hello world! 3")
		logrus.Info("Hello world! 4")
		logrus.Info("Hello world! 5")
		logrus.Info("Hello world! 6")
		logrus.Info("Hello world! 7")
		logrus.Warn("Hello world! 8")

		logs := buffer.Read()
		for _, l := range logs {
			fmt.Println(l.Timestamp, l.Level, l.Message)
		}
	})
}

func TestLogrusBufferHook_ThreadSafe(t *testing.T) {
	Convey("logrus buffer hook", t, func() {
		hook, buffer := NewLogrusBufferHook([]logrus.Level{logrus.WarnLevel, logrus.InfoLevel}, 10)

		logrus.AddHook(hook)
		logrus.SetOutput(ioutil.Discard)

		go func() {
			for i := 0; i < 10000; i++ {
				logrus.Infof("Hello world! From-1 %d", i)
				time.Sleep(time.Millisecond * 50)
			}
		}()

		go func() {
			for i := 0; i < 10000; i++ {
				logrus.Infof("Hello world! From-2 %d", i)
				time.Sleep(time.Millisecond * 100)
			}
		}()

		go func() {
			for i := 0; i < 20; i++ {
				logs := buffer.Read()
				for _, l := range logs {
					fmt.Println(l)
				}
				time.Sleep(time.Second)
			}
		}()

		time.Sleep(time.Second * 30)
	})
}
