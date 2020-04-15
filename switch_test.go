package main

import (
	"fmt"
	"testing"
)

func TestSwitch(t *testing.T) {
	switch 100 {
	case 1, 2, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14:
		fmt.Println("employ")
	case 3, 4:
		fmt.Println("leader")
	default:
		fmt.Println("default")
	}
}

func TestSwitch1(t *testing.T) {
	for i := 0; i < 10; i++ {
		switch i {
		case 1,
			2,
			3,
			4,
			5:
			fmt.Println("normal ", i)
		case 6,
			7:
			continue
			fmt.Println("continue nothing ", i)
		default:
			fmt.Println("default ", i)
		}
	}
}
