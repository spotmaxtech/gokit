package gokit

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestLogrusFormatter(t *testing.T) {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(FileLineFormatter)
	logrus.Infof("hello world!")
}
