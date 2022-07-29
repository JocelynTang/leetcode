package golang

func reverse(x int) int {
	INT_MAX := 2147483647
	INT_MIN := -INT_MAX - 1
	newX := 0
	if x == 0 {
		return 0
	}
	for {
		newX = newX*10 + x%10
		x = x / 10
		if x == 0 {
			break
		}
	}
	if newX > INT_MAX || newX < INT_MIN {
		return 0
	}
	return newX
}
