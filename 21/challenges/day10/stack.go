package day10

import (
	"errors"
)

type stack []rune

var emptyStackError = errors.New("Stack is empty")

func (s *stack) push(i rune) {
	*s = append(*s, i)
}

func (s *stack) pop() (rune, error) {
	var popped rune
	if len(*s) == 0 {
		return -1, emptyStackError
	}
	popped = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return popped, nil
}

func (s *stack) peek() rune {
	if s.isEmpty() {
		return -1
	} else {
		return (*s)[len(*s)-1]
	}
}

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}
