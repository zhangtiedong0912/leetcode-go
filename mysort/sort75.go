package mysort

func SortColors(nums []int) {
	if len(nums) <= 1 {
		return
	}
	//三指针
	var left = 0
	var right = len(nums) - 1
	var curr = 0
	for curr <= right {
		// 0 left与curr互换 left++ curr++
		if nums[curr] == 0 {
			nums[curr], nums[left] = nums[left], nums[curr]
			left++
			curr++

			// 2 right与curr互换 right-- 但是curr不能加，因为不知道换过来的是啥(上面换过来的结果是确定的)
		} else if nums[curr] == 2 {
			nums[curr], nums[right] = nums[right], nums[curr]
			right--
			// 1 不需要换， curr++
		} else {
			curr++
		}
	}
}
