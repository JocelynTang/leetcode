package golang

import (
	"sort"
)

// fail
func ThreeSumClosest(nums []int, target int) int {
	n := len(nums)
	sort.Ints(nums)
	var ans int
	var (
		abs = func(x int) int {
			if x < 0 {
				return -x
			}
			return x
		}
	)

	closest := abs(target - (nums[0] + nums[1] + nums[2]))
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		k := n - 1
		for j := i + 1; j < n; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for j < k && abs(nums[i]+nums[j]+nums[k]) > abs(target-closest) {
				k--
			}
			if j == k {
				break
			}
			if abs(target-nums[i]-nums[j]-nums[k]) <= closest {
				ans = nums[i] + nums[j] + nums[k]
				closest = abs(target - ans)
			}
		}
	}
	return ans
}

/*
func threeSumClosest(nums []int, target int) int {
    n := len(nums)
	sort.Ints(nums)
	var ans int
	var (
		abs = func(x int) int {
			if x < 0 {
				return -x
			}
			return x
		}
	)

	closest := abs(target - (nums[0] + nums[1] + nums[2]))
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				sum := nums[i] + nums[j] + nums[k]
				if abs(target-sum) <= closest {
					ans = sum
					closest = abs(target - sum)
				}
			}
		}
	}
	return ans
}
*/
