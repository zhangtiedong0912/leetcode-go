package stack

/*
输入: [1,2,1]
输出: [2,-1,2]
解释: 第一个 1 的下一个更大的数是 2；
数字 2 找不到下一个更大的数；
第二个 1 的下一个最大的数需要循环搜索，结果也是 2。
*/
func nextGreaterElements(nums []int) []int {
	len := len(nums)
	res := make([]int,len)
	stack := InitMyStack()
	for i:= 2 * len - 1;i>=0;i-- {
		for !stack.IsEmpty() && stack.Peek().(int) <= nums[i%len] {
			stack.Pop()
		}
		if stack.IsEmpty(){
			res[i%len] = -1
		}else{
			res[i%len] = stack.Peek().(int)
		}
		stack.Push(nums[i%len])
	}

	return res
}
