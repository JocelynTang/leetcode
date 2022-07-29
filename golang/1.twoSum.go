package golang

func twoSum(nums []int, target int) []int {
	var twoNum []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				twoNum = append(twoNum, i, j)
			}
		}
	}
	return twoNum
}
