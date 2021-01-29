package gokit

import (
	"fmt"
	"testing"
)

func TestRingbuffer(t *testing.T) {
	buf := NewRingBuffer(5)

	// buf.Write(&RingBufferData{Message: "a", Level: "info"})
	// buf.Write(&RingBufferData{Message: "b", Level: "info"})
	// buf.Write(&RingBufferData{Message: "c", Level: "info"})
	// buf.Write(&RingBufferData{Message: "d", Level: "info"})
	// buf.Write(&RingBufferData{Message: "e", Level: "info"})
	buf.Write(&RingBufferData{Message: "f", Level: "warn"})
	buf.Write(&RingBufferData{Message: "g", Level: "info"})
	buf.Write(&RingBufferData{Message: "h", Level: "info"})
	fmt.Println(PrettifyJson(buf.Read(), false))
	fmt.Println(PrettifyJson(buf.Read(), false))
}
