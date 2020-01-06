package stack

/*
输入：[3,1,2,4]
输出：17

1224

[3]       stack 3,1  tmp 3              res 3
[3,1]     stack 1,2  tmp 3-3*1+2*1      res 5
[3,1,2]   stack 2,1  tmp 2+2            res 9
[3,1,2,4] stack 4,1  tmp 2+2+4          res 17
*/
type data1 struct{
	Index int
	Count int
}
func sumSubarrayMins(A []int) int {
	res := 0
	tmp := 0
	m := 1000000007
	stack := InitMyStack()

	for _,v := range A {
		count := 1
		for !stack.IsEmpty() && stack.Peek().(data1).Index >= v {
			data := stack.Pop().(data1)
			count += data.Count
			tmp -= data.Index * data.Count
		}
		stack.Push(data1{Index:v,Count:count})
		tmp += v * count
		res += tmp
		res %=m
	}


	return res
}
