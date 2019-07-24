package gokit

import "time"

type SimpleTimer struct {
	duration time.Duration
	last     time.Time
}

func NewSimpleTimer(duration time.Duration) *SimpleTimer {
	timer := &SimpleTimer{
		duration: duration,
		last:     time.Now(),
	}
	return timer
}

// if time duration past
func (t *SimpleTimer) Timeout() bool {
	return time.Since(t.last) > t.duration
}

// reset timer to now
func (t *SimpleTimer) Reset() {
	t.last = time.Now()
}

// if timeout, reset it automatically
func (t *SimpleTimer) Checkpoint() bool {
	if t.Timeout() {
		t.Reset()
		return true
	} else {
		return false
	}
}
