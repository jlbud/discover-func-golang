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

func TestSwitch2(t *testing.T) {
	for i := 0; i < 50; i++ {
		if i >= 4 && i <= 14 || i == 21 {
			t.Log("eligible:", i)
		}
	}

	t.Log("success")
}

func TestSwitch3(t *testing.T) {
	var ScoreCoefficient float32
	ScoreCoefficient = 1
	switch "enWordRead" {
	case "enWordRead",
		"enSentRead",
		"enParaRead",
		"enPhonics",
		"enSentImit",
		"enParaImit",
		"enChoice",
		"enAskAnswer",
		"enTopic",
		"enPict",
		"enRetell":
		if ScoreCoefficient == 0 || ScoreCoefficient < 0.6 || ScoreCoefficient > 1.9 {
			fmt.Println("ScoreCoefficient", ScoreCoefficient)
		}
	}
	fmt.Println("end...")
}

func TestSwitch4(t *testing.T) {
	type A1 struct {
		a string
	}
	type A2 struct {
		a string
	}
	var v interface{}
	v = &A1{}
	//v = A1{}
	switch v.(type) {
	case *A1:
		fmt.Println("A1")
	case *A2:
		fmt.Println("A2")
	default:
		fmt.Println("default")
	}
}
