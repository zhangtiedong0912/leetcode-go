package stack



type stack []int

func (s *stack) Push(v int){
	*s = append((*s), v)
}

func (s *stack) Pop() int{
	l := len(*s)
	v := (*s)[l - 1]
	*s = (*s)[:l - 1]

	return v
}

func (s *stack) Seek() int{
	l := len(*s)
	v := (*s)[l - 1]

	return v
}

func (s *stack) Len() int{
	return len(*s)
}

const (
	INT_MAX = int(^uint(0) >> 1)
	INT_MIN = ^INT_MAX

)

//3, 1, 4, 2
/*
思路：利用单调栈性质，从后往前，找到小于次大的值即可

*/
func find132pattern(nums []int) bool {

	if len(nums) < 3 {
		return false
	}
	stack := stack{}
	len :=len(nums)-1

	second := INT_MIN
	for i:=len;i>=0;i--{
		if nums[i] < second{
			return true
		}

		for stack.Len()!= 0 && nums[i]>stack.Seek(){
			second = max(second, stack.Pop())
		}
		stack.Push(nums[i])

	}

	return false
}

func max(i,j int) int {
	if i>j{
		return i
	}
	return j
}