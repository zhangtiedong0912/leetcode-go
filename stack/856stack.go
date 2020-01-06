package stack

//()()
func scoreOfParentheses(S string) int {
	stack := []int{}
	stack = append(stack, 0)

	for _, v := range S {
		if byte(v) == '(' {
			stack = append(stack, 0)
		} else {
			v := stack[len(stack)-1]
			w := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, w+max(v*2, 1))
		}
	}

	return stack[len(stack)-1]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
