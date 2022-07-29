package golang

import "strconv"

// 9
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x > 0) {
		return false
	}
	test := x
	newX := 0
	for i := 0; i < len(strconv.Itoa(x)); i++ {
		newX = newX*10 + test%10
		test = test / 10
	}
	return newX == x
}
