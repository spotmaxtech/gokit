package gokit

// derived from github, customized and support v7

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"strings"
	"time"
)

type ESIndexRotate int

const (
	RotateNon ESIndexRotate = iota
	RotateDay
	RotateMonth
)

type ElasticHook struct {
	client    *elastic.Client
	host      string
	index     string
	rotate    ESIndexRotate
	levels    []logrus.Level
	ctx       context.Context
	ctxCancel context.CancelFunc
}

//
func NewElasticHook(client *elastic.Client, host string, level logrus.Level, index string, rotate ESIndexRotate) (*ElasticHook, error) {
	var levels []logrus.Level
	for _, l := range []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	} {
		if l <= level {
			levels = append(levels, l)
		}
	}

	ctx, cancel := context.WithCancel(context.TODO())

	return &ElasticHook{
		client:    client,
		host:      host,
		index:     index,
		rotate:    rotate,
		levels:    levels,
		ctx:       ctx,
		ctxCancel: cancel,
	}, nil
}

func (hook *ElasticHook) Fire(entry *logrus.Entry) error {
	level := entry.Level.String()

	if errData, ok := entry.Data[logrus.ErrorKey]; ok {
		if err, ok := errData.(error); ok && entry.Data[logrus.ErrorKey] != nil {
			entry.Data[logrus.ErrorKey] = err.Error()
		}
	}

	msg := struct {
		Host      string
		Timestamp string
		Message   string
		Data      logrus.Fields
		Level     string
	}{
		hook.host,
		entry.Time.UTC().Format(time.RFC3339Nano),
		entry.Message,
		entry.Data,
		strings.ToUpper(level),
	}

	// rotate the index
	var index string
	switch hook.rotate {
	case RotateDay:
		index = fmt.Sprintf("%s.%s", hook.index, entry.Time.UTC().Format("20060101"))
	case RotateMonth:
		index = fmt.Sprintf("%s.%s", hook.index, entry.Time.UTC().Format("200601"))
	case RotateNon:
		index = hook.index
	default:
		index = hook.index
	}

	_, err := hook.client.
		Index().
		Index(index).
		// Type("log").
		BodyJson(msg).
		Do(hook.ctx)
	return err
}

func (hook *ElasticHook) Levels() []logrus.Level {
	return hook.levels
}

func (hook *ElasticHook) Cancel() {
	hook.ctxCancel()
}
