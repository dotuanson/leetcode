package main

import "math"

func maxSubArray(nums []int) int {
	var lengthNums = len(nums)
	var largestSums = make([]int, lengthNums)
	largestSums[0] = nums[0]

	for idx, num := range nums {
		if idx == 0 {
			continue
		}
		largestSums[idx] = int(math.Max(float64(largestSums[idx-1]+num), float64(num)))
	}

	result := largestSums[0]
	for idx, val := range largestSums {
		if idx == 0 {
			continue
		}
		result = int(math.Max(float64(result), float64(val)))
	}
	return result
}
