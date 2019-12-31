package mysort

import "sort"

func Merge1666(intervals [][]int) int {
	var res = 1
	if len(intervals) <= 1 {
		return res
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] <= intervals[j][0]
	})

	var left = intervals[0][0]
	var right = intervals[0][1]

	for i := 1; i < len(intervals); i++ {

		if intervals[i][0] > right {
			if i == 1 {
				res = res + 2
			} else {
				res++
			}
			left = intervals[i][0]
			right = intervals[i][1]
		} else if intervals[i][0] <= left && intervals[i][1] <= right {
			left = intervals[i][1]
		} else if intervals[i][0] <= left && intervals[i][1] >= right {
			left = right
			right = intervals[i][1]
		} else if intervals[i][0] > left && intervals[i][1] <= right {
			res++
			left = intervals[i][1]
		} else if intervals[i][0] > left && intervals[i][1] >= right {
			res++
			left = right
			right = intervals[i][1]
		}
	}
	return res
}
