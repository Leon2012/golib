package ds

import (
	"errors"
	"fmt"
	"sync"
)

type stacknode struct {
	data interface{}
	next *stacknode
}

type Stack struct {
	head  *stacknode
	count int
	lock  *sync.Mutex
}

func NewStack() *Stack {
	s := &Stack{}
	s.lock = &sync.Mutex{}
	return s
}

// Get() peeks at the n-th item in the stack. Unlike other operations, this one costs O(n).
func (s *Stack) Get(index int) (interface{}, error) {
	if index < 0 || index >= s.count {
		return nil, errors.New(fmt.Sprintf("Requested index %d outside stack, length %d", index, s.count))
	}

	s.lock.Lock()
	defer s.lock.Unlock()

	n := s.head
	for i := 1; i < s.count-index; i++ {
		n = n.next
	}

	return n.data, nil
}

// Dump() prints a textual representation of the stack.
func (s *Stack) Dump() {
	n := s.head
	fmt.Print("[ ")
	for i := 0; i < s.count; i++ {
		fmt.Printf("%+v ", n.data)
		n = n.next
	}
	fmt.Print("]")
}

func (s *Stack) Len() int {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.count
}

func (s *Stack) Push(item interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	n := &stacknode{data: item}

	if s.head == nil {
		s.head = n
	} else {
		n.next = s.head
		s.head = n
	}

	s.count++
}

func (s *Stack) Pop() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()

	var n *stacknode
	if s.head != nil {
		n = s.head
		s.head = n.next
		s.count--
	}

	if n == nil {
		return nil
	}

	return n.data

}

func (s *Stack) Peek() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()

	n := s.head
	if n == nil || n.data == nil {
		return nil
	}

	return n.data
}
