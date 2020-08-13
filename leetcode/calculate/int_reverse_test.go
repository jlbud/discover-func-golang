package calculate

import (
	"fmt"
	"math"
	"testing"
)

func TestReverse(t *testing.T) {
	i := reverse(1190)
	fmt.Println(i)
}

/**
整数反转
*/
func reverse(x int) int {
	ret := 0
	for x != 0 {
		// 拿到原整数串中最后一个数
		last := x % 10
		// 原整数串去掉最后一位数，利用了向上取整
		//x /= 10
		x = x / 10
		// 最后一位数乘以10，扩充一位
		ret = ret*10 + last
		if ret < math.MinInt32 || ret > math.MaxInt32 {
			return 0
		}
	}
	return ret
}

func TestIsPalindrome(t *testing.T) {
	palindrome := isPalindrome(1)
	fmt.Println(palindrome)
}

/**
整数是否是回文数
*/
func isPalindrome(x int) bool {
	// 如果是负数肯定不是
	if x < 0 {
		return false
	}
	// 缓存原数
	cache := x
	// 反转原数
	ret := 0
	for x != 0 {
		last := x % 10
		// x /= 10
		x = x / 10
		ret = ret*10 + last
		if ret < math.MinInt32 || ret > math.MaxInt32 {
			return false
		}
	}
	// 反转的数和原数字相等，则是回文数
	if ret == cache {
		return true
	}
	return false
}
