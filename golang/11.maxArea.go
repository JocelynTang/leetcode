package golang

func maxArea(height []int) int {
	n := len(height)
	var maxarea, h, area int
	/*
		//timeout
			for i := 0; i < n-1; i++ {
				for j := n - 1; j > i; j-- {
					if height[i] < height[j] {
						h = height[i]
					} else {
						h = height[j]
					}
					area = (j - i) * h
					if area > maxarea {
						maxarea = area
					}
				}
			}
	*/
	for i, j := 0, n-1; i < j; {

		if height[i] < height[j] {
			h = height[i]
		} else {
			h = height[j]
		}
		area = (j - i) * h
		if area > maxarea {
			maxarea = area
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return maxarea
}
