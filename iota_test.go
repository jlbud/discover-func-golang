package main

import (
	"fmt"
	"testing"
)

const (
	readyStartState = iota + 1
	collectingState
	completeState
	suspendedState
	endState
	deleteState
)

func operate(inputState int) {
	switch inputState {
	case readyStartState:
		fmt.Println(readyStartState)
	case collectingState:
		fmt.Println(collectingState)
	case completeState:
		fmt.Println(completeState)
	case suspendedState:
		fmt.Println(suspendedState)
	case endState:
		fmt.Println(endState)
	case deleteState:
		fmt.Println(deleteState)
	}
}

func TestOutputState(t *testing.T) {
	fmt.Println("readyStartState", readyStartState)
	fmt.Println("deleteState", deleteState)
}

func TestOperateState(t *testing.T) {
	operate(1)
}
