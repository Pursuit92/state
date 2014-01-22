package main

import (
	"github.com/Pursuit92/state"
	"fmt"
)

type embeddedState struct {
	st state.Simple
	transitions int
}

func newES() *embeddedState {
	es := embeddedState{}
	return &es
}

func (e *embeddedState) Pop() (state.StateTrans,error) {
	return e.st.Pop()
}

func (e *embeddedState) Push(states ...state.StateTrans) {
	e.st.Push(states...)
}

func (e *embeddedState) Len() int {
	return e.st.Len()
}

func state1(s state.StateMachine,i interface{}) error {
	es := s.(*embeddedState)
	es.transitions++
	fmt.Println("State 1!")
	s.Push(state2)
	return nil
}

func state2(s state.StateMachine,i interface{}) error {
	es := s.(*embeddedState)
	es.transitions++
	fmt.Println("State 2!")
	return nil
}

func main() {
	s := newES()
	s.Push(state1,state2,state2)
	for i := 0; i < 10; i++ {
		err := state.Input(s,i)
		fmt.Println("machine is",s.Len(),"states long")
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("Transitioned",s.transitions,"times!")
}

