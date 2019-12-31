package mysort

/*
思路：先构建一个大小为K的小顶堆，然后遍历剩余元素，若大于堆顶，与第一个交换，维护堆的平衡
最种是一个存储了最大K个数的小顶堆，则堆顶也就是第K大的数
*/
func findKthLargest(nums []int, k int) int {
	heap := nums[:k]
	//构造一个大小为K的小顶堆
	buildHeap(heap)

	//遍历剩下的元素，若大于堆顶，与第一个交换，维护堆的平衡
	for i := k; i < len(nums); i++ {
		if nums[i] > heap[0] {
			heap[0] = nums[i]
			heapify(heap, 0)
		}
	}
	return heap[0]
}

func buildHeap(nums []int) {
	n := len(nums)
	//从最后一个叶子节点的父节点开始向上维护平衡
	for i := n / 2; i >= 0; i-- {
		heapify(nums, i)
	}
}

//维护堆的平衡
func heapify(nums []int, index int) {
	n := len(nums)
	for index < n {
		//记录当前节点,以及当前节点的左右孩子
		var min = index
		var left = index*2 + 1
		var right = index*2 + 2

		//寻找子节点中最小的那个，如果比父小，交换
		if left < n && nums[left] < nums[min] {
			min = left
		}
		if right < n && nums[right] < nums[min] {
			min = right
		}

		//找到比父小的，交换，挪动index (维持上层的平衡）
		if min != index {
			nums[min], nums[index] = nums[index], nums[min]
			index = min
		} else {
			break
		}
	}
}
