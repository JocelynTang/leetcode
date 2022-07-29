package golang

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	newNums := make([]int, len(nums1)+len(nums2))
	var median float64
	i, j, k := 0, 0, 0
	for k < len(newNums) {
		for i < len(nums1) && j < len(nums2) {
			if nums1[i] < nums2[j] {
				newNums[k] = nums1[i]
				i++
			} else {
				newNums[k] = nums2[j]
				j++
			}
			k++
		}
		if i >= len(nums1) && j < len(nums2) {
			for j < len(nums2) {
				newNums[k] = nums2[j]
				k++
				j++
			}
		}
		if j >= len(nums2) && i < len(nums1) {
			for i < len(nums1) {
				newNums[k] = nums1[i]
				k++
				i++
			}
		}
	}

	half := len(newNums) / 2
	if len(newNums)%2 == 0 {
		median = (float64(newNums[half]) + float64(newNums[half-1])) / 2
	} else {
		median = float64(newNums[half])
	}
	return median
}
