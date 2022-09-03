package stack

import (
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
	if g, _ := s.Peek(); g != 10 {
		t.Fatal("should have returned ten")
	}
	s.Pop()
	s.Peek()
}

func TestStackPush(t *testing.T) {
	s := newStack(1)
	s.Push(10)
	s.Push(11)
	s.Push(1)
	s.Push(10)
}

func TestStackPop(t *testing.T) {
	s := newStack(4)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(12)

	if g, _ := s.Pop(); g != 12 {
		t.Fatal("should have returned twelve")
	}
	if g, _ := s.Pop(); g != 3 {
		t.Fatal("should have returned three")
	}
	if g, _ := s.Pop(); g != 2 {
		t.Fatal("should have returned two")
	}
	if g, _ := s.Pop(); g != 1 {
		t.Fatal("should have returned one")
	}
	if _, err := s.Pop(); err == nil {
		t.Fatal("should have returned an error")
	}
}
