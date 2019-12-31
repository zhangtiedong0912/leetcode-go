package mysort

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	//排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	//合并
	var result [][]int
	result = append(result, intervals[0])
	for i := 1; i < len(intervals); i++ {
		var a = len(result) - 1
		var b = len(result[a]) - 1

		if result[a][b] >= intervals[i][0] {
			var j = len(intervals[i]) - 1
			result[a][b] = max(result[a][b], intervals[i][j])

		} else {
			result = append(result, intervals[i])
		}
	}
	return result
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
