package main

type Stack []uint

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) push(num uint) {
	*s = append(*s, num)
}

func (s *Stack) pop() (uint, bool) {

	if s.isEmpty() {
		return 0, false
	}

	lastIndex := len(*s) - 1
	removedElement := (*s)[lastIndex]
	*s = (*s)[:lastIndex]
	return removedElement, true
}
