package gokit

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"path"
	"runtime"
)

var FileLineFormatter = &logrus.TextFormatter{
	CallerPrettyfier: func(f *runtime.Frame) (string, string) {
		filename := path.Base(f.File)
		return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
	},
}
