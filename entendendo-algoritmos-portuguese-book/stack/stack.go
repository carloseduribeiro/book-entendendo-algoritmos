package stack

import (
	//"fmt"
	"errors"
)

type node struct {
	value int
	tail *node
}

type stack struct {
	top *node
}

func (s *stack) push(v int) {
	if s.top == nil {
		s.top = &node{ value: v }
	}
	//tmp := *s.top
	s.top = &node{value: v, tail: s.top}
}

// todo - fix pop
func (s *stack) pop() (int, error) {
	if s.top != nil {
		value := s.top.value
		s.top = s.top.tail
		return value, nil
	}
	return 0, errors.New("stask is empty")
}

//func main() {
//	s := stack{}
//	s.push(3)
//	s.push(2)
//	s.push(1)
//
//	pop1, err := s.pop()
//	if err != nil {
//		panic(err)
//	}
//	fmt.Printf("stash pop() = %d\n", pop1)
//
//	pop2, err := s.pop()
//	if err != nil {
//		panic(err)
//	}
//	fmt.Printf("stash pop() = %d\n", pop2)
//
//	pop3, err := s.pop()
//	if err != nil {
//		panic(err)
//	}
//	fmt.Printf("stash pop() = %d\n", pop3)
//
//	pop4, err := s.pop()
//	if err != nil {
//		panic(err)
//	}
//	fmt.Printf("stash pop() = %d\n", pop4)
//}
