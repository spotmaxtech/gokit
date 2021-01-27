package gokit

import (
	"fmt"
	"testing"
)

func TestRingbuffer(t *testing.T) {
	buf := NewRingBuffer(5)
	buf.Write("a")
	buf.Write("b")
	buf.Write("c")
	buf.Write("d")
	buf.Write("e")
	buf.Write("f")
	buf.Write("g")
	buf.Write("h")
	buf.Write("j")
	buf.Write("k")
	fmt.Println(buf.Read())
	fmt.Println(buf.Read())
}
