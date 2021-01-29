package gokit

import (
	"sync"
)

type RingBufferData struct {
	Timestamp string `json:"timestamp"`
	Message   string `json:"message"`
	Level     string `json:"level"`
}

type RingBuffer struct {
	buf  []*RingBufferData
	size int
	w    int // next position to write
	mu   sync.Mutex
}

func NewRingBuffer(size int) *RingBuffer {
	if size <= 0 {
		panic("buffer size must larger than 0 !!!")
	}
	return &RingBuffer{
		buf:  make([]*RingBufferData, size),
		size: size,
	}
}

func (r *RingBuffer) Write(p *RingBufferData) {

	r.mu.Lock()
	r.buf[r.w] = p
	r.w += 1
	if r.w == r.size {
		r.w = 0
	}
	r.mu.Unlock()

	return
}

func (r *RingBuffer) Read() []*RingBufferData {
	r.mu.Lock()
	var data []*RingBufferData
	for i, count := r.w-1, 0; count < r.size; {
		if i < 0 {
			i = r.size - 1
		}
		if r.buf[i] != nil {
			data = append(data, r.buf[i])
		} else {
			break
		}
		i--
		count++
	}

	r.mu.Unlock()
	return data
}
