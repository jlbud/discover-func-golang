package leetcode

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	// 下标为0,1,2,3
	nums := []int{2, 11, 15, 7}
	ints := twoSum(nums, 9)
	// 返回[0 3]
	fmt.Println(ints)
}

/*
	给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标
*/
func twoSum(nums []int, target int) []int {
	// 先遍历数组的每一个数
	for i, v := range nums {
		// 再遍历这个数组大于当前这个数的每一个数
		for k := i + 1; k < len(nums); k++ {
			// 如果目标数减去第一个遍历的数正好等于当前遍历的数，
			// 则直接返回下标
			if target-v == nums[k] {
				return []int{i, k}
			}
		}
	}
	return []int{}
}
