package state

import (
	"reflect"
	"fmt"
	"container/list"
)

type StateTrans func(*State,interface{}) error
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

func (s *State) transOne(i int) error {
	return nil
}

type State struct {
	list.List
	Final bool
}

func (s *State) Input(i interface{}) error {
	front := s.Front()
	if front != nil {
		f,ok := front.Value.(func(*State,interface{})error)
		if ok {
			s.Remove(front)
			return f(s,i)
		}
		return stateErr{Message: InvalidTrans,Info: fmt.Sprintf("%v",reflect.TypeOf(front.Value))}
	}
	return stateErr{Message: StackEmpty}
}
