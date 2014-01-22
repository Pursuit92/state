package state

import (
	"reflect"
	"fmt"
	"container/list"
)

type StateTrans func(StateMachine,interface{}) error
type stateErr struct {
	Message, Info string
}

func (s stateErr) Error() string {
	if s.Info != "" {
		return fmt.Sprintf("%s: %s",s.Message,s.Info)
	} else {
		return s.Message
	}
}

const (
	StackEmpty = "State Stack is empty!"
	InvalidTrans = "Invalid Transition!"
)

func (s *Simple) transOne(i int) error {
	return nil
}

type StateMachine interface {
	Push(...StateTrans)
	Pop() (StateTrans,error)
}

func Input(s StateMachine, i interface{}) error {
	f,err := s.Pop()
	if err == nil {
		return f(s,i)
	}
	return err
}

type Simple struct {
	stack list.List
}

func New() *Simple {
	return &Simple{}
}

func (s *Simple) Push(states ...StateTrans) {
	for i := len(states) - 1; i >= 0; i-- {
		s.stack.PushFront(states[i])
	}
}

func (s *Simple) Pop() (StateTrans, error) {
	if s.stack.Len() != 0 {
		front := s.stack.Front()
		f, ok := front.Value.(StateTrans)
		if ok {
			s.stack.Remove(front)
			return f,nil
		}
		return nil,stateErr{Message: InvalidTrans, Info: fmt.Sprintf("%v",reflect.TypeOf(front.Value))}
	}
	return nil, stateErr{Message: StackEmpty}
}

func (s *Simple) Len() int {
	return s.stack.Len()
}
