package gokit

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"strings"
	"time"
)

type LogrusBufferHook struct {
	levels    []logrus.Level
	buffer    *RingBuffer
	ctx       context.Context
	ctxCancel context.CancelFunc
}

// return
// hook: for logrus
// buffer: for reading data
func NewLogrusBufferHook(levels []logrus.Level, size int) (*LogrusBufferHook, *RingBuffer) {
	ctx, cancel := context.WithCancel(context.TODO())

	buffer := NewRingBuffer(size)
	return &LogrusBufferHook{
		levels:    levels,
		ctx:       ctx,
		ctxCancel: cancel,
		buffer:    buffer,
	}, buffer
}

func (hook *LogrusBufferHook) Fire(entry *logrus.Entry) error {
	level := entry.Level.String()

	if errData, ok := entry.Data[logrus.ErrorKey]; ok {
		if err, ok := errData.(error); ok && entry.Data[logrus.ErrorKey] != nil {
			entry.Data[logrus.ErrorKey] = err.Error()
		}
	}

	msg := &RingBufferData{
		entry.Time.UTC().Format(time.RFC3339Nano),
		entry.Message,
		strings.ToUpper(level),
	}

	hook.buffer.Write(msg)
	return nil
}

func (hook *LogrusBufferHook) Levels() []logrus.Level {
	return hook.levels
}

func (hook *LogrusBufferHook) Cancel() {
	hook.ctxCancel()
}
