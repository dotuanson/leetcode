package main

func twoSum(nums []int, target int) []int {
	for indexOne, valueOne := range nums {
		for indexTwo, valueTwo := range nums {
			if indexOne != indexTwo {
				if valueOne+valueTwo == target {
					return []int{indexOne, indexTwo}
				}
			}
		}
	}
	return []int{}
}
