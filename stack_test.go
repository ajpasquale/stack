package stack

import (
	"fmt"
	"testing"
)

func TestNewStack(t *testing.T) {
	s := newStack(0)
	if s == nil {
		t.Fatal("should return new stack")
	}
}

func TestStackPeek(t *testing.T) {
	s := newStack(1)
	s.Push(10)
	fmt.Println(s.Peek())
}
