package main

import (
	"testing"
)

func TestStack_isEmpty(t *testing.T) {
	emptyStack := Stack{}

	tests := []struct {
		name string
		s    *Stack
		want bool
	}{
		{"should return true", &emptyStack, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.isEmpty(); got != tt.want {
				t.Errorf("Stack.isEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_push(t *testing.T) {
	stack := Stack{1, 2}
	stack.push(3)

	if len(stack) != 3 {
		t.Errorf("Stack size is '%d'", len(stack))
	}
}

func TestStack_pop(t *testing.T) {
	stack := Stack{1, 2, 3}

	expected := stack[len(stack)-1]

	popped, ok := stack.pop()

	if !ok || expected != popped {
		t.Errorf("Expected '%v' but '%v' was popped.", expected, popped)
	}
}
