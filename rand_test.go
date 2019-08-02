package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRand(t *testing.T) {
	rand.Seed(time.Now().Unix())
	questionIds := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
	perm := rand.Perm(len(questionIds))

	fmt.Println(perm[:5])
}
