package main

import (
	"fmt"
	"testing"
)

func split(arr []int, low int, high int, val int) int {
	var mid = (low + high) / 2
	if low <= high {
		if arr[mid] == val {
			return mid
		} else {
			if arr[mid] < val {
				return split(arr, mid+1, high, val)
			} else {
				return split(arr, low, mid-1, val)
			}
		}
	} else {
		fmt.Println(low)
		return low
	}
}

// 树转二叉树
func TestSplit(t *testing.T) {
	arr := []int{1, 4, 6, 12, 17, 19, 44, 66, 71}
	t.Log(split(arr, 0, len(arr)-1, 44))
	t.Log(split(arr, 0, len(arr)-1, 55))
}
