package main

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	var newIntervals [][]int
	var isIntersection bool
	for _, interval := range intervals {
		if (interval[0] >= newInterval[0]) && (!isIntersection) {
			newIntervals = append(newIntervals, newInterval)
			isIntersection = true
		}
		newIntervals = append(newIntervals, interval)
	}
	if !isIntersection {
		newIntervals = append(newIntervals, newInterval)
	}
	result := [][]int{newIntervals[0]}
	for idx, interval := range newIntervals {
		if idx == 0 {
			continue
		}
		if result[len(result)-1][1] < interval[0] {
			result = append(result, interval)
		} else {
			if interval[1] > result[len(result)-1][1] {
				result[len(result)-1][1] = interval[1]
			}
		}
	}
	return result
}
