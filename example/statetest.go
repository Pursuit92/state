package main

import (
	"github.com/Pursuit92/state"
	"fmt"
)

func state1(s *state.State,i interface{}) error {
	fmt.Println("State 1!")
	s.PushFront(state2)
	return nil
}

func state2(s *state.State,i interface{}) error {
	fmt.Println("State 2!")
	s.PushFront(state1)
	return nil
}

func main() {
	s := &state.State{}
	s.PushFront(state1)
	for i := 0; i < 10; i++ {
		err := s.Input(i)
		fmt.Println("machine is",s.Len(),"states long")
		if err != nil {
			fmt.Println(err)
		}
	}
}

