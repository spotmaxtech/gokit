package gokit

import (
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestElasticHook(t *testing.T) {
	Convey("logrus elasticsearch hook", t, func() {
		log := logrus.New()
		client, err := elastic.NewClient(elastic.SetURL("http://es.spotmaxtech.com"),
			elastic.SetBasicAuth("", ""),
			elastic.SetSniff(false))
		So(err, ShouldEqual, nil)

		hook, err := NewElasticHook(client, "localhost", logrus.WarnLevel, "testlog", RotateDay)
		So(err, ShouldBeNil)

		log.AddHook(hook)

		log.WithFields(logrus.Fields{
			"name": "spotmax",
			"age":  15,
		}).Warn("Hello world!")

		log.WithFields(logrus.Fields{
			"name": "spotmax",
			"age":  100,
		}).Info("Hello world!")
	})
}
