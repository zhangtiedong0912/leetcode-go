package stack

//[73, 74, 75, 71, 69, 72, 76, 73]
//[1, 1, 4, 2, 1, 1, 0, 0]
type myData struct {
	Index int
	Value int
}

func dailyTemperatures(T []int) []int {
	res := make([]int, len(T))
	stack := InitMyStack()

	for i := len(T) - 1; i >= 0; i-- {
		for !stack.IsEmpty() && T[i] >= stack.Peek().(myData).Value {
			stack.Pop()
		}
		if stack.IsEmpty() {
			res[i] = 0
		} else {
			res[i] = stack.Peek().(myData).Index - i
		}
		stack.Push(myData{Index: i, Value: T[i]})
	}
	return res
}
