package main

import (
	"fmt"
	"reflect"

	lt "leetcode/golang"
)

func main() {
	strs := []string{"flower", "flow", "flight"}
	pre := lt.LongestCommonPrefix(strs)
	fmt.Printf("the common prefix is %s\n", pre)

	//nums := []int{-2, -1, 0, 1, 2, 3}
	//nums := []int{0, 0, 0, 0, 0, 0}
	nums := []int{0, 1, 2}
	threeNums := lt.ThreeSum(nums)
	fmt.Printf("the nums are: %v\n", threeNums)
	fmt.Printf("the type of nums is: %v\n", reflect.TypeOf(threeNums))

	threeSumClosest := lt.ThreeSumClosest(nums, 3)
	fmt.Printf("the closest value is: %d\n", threeSumClosest)

}
