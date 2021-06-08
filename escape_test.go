package main

// remove escape character
// https://github.com/jialanli/lacia/blob/main/utils/str_test.go
func TextRemoveX(str string, x string) string {
	var res string
	for i := 0; i < len(str); i++ {
		if string(str[i]) != x {
			res = res + string(str[i])
		}
	}
	return res
}
