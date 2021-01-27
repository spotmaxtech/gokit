package gokit

import (
	"sync"
)

type RingBuffer struct {
	buf  []string
	size int
	w    int // next position to write
	mu   sync.Mutex
}

func NewRingBuffer(size int) *RingBuffer {
	if size <= 0 {
		panic("buffer size must larger than 0 !!!")
	}
	return &RingBuffer{
		buf:  make([]string, size),
		size: size,
	}
}

func (r *RingBuffer) Write(p string) {
	if len(p) == 0 {
		return
	}

	r.mu.Lock()
	r.buf[r.w] = p
	r.w += 1
	if r.w == r.size {
		r.w = 0
	}
	r.mu.Unlock()

	return
}

func (r *RingBuffer) Read() []string {
	r.mu.Lock()
	var data []string
	for i, count := r.w-1, 0; count < r.size; {
		if i < 0 {
			i = r.size - 1
		}
		if r.buf[i] != "" {
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
