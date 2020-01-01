package stack

func backspaceCompare(S string, T string) bool {
	stack1 := []int32{}
	stack2 := []int32{}

	for _,v := range S {
		if byte(v) == '#' {
			if len(stack1)!= 0 {
				stack1 = stack1[:len(stack1)-1]
			}
		}else {
			stack1 = append(stack1, v)
		}
	}

	for _,v := range T {
		if byte(v) == '#' {
			if len(stack2)!= 0 {
				stack2 = stack2[:len(stack2)-1]
			}
		}else {
			stack2 = append(stack2, v)
		}
	}

	if len(stack1) != len(stack2) {
		return false
	}
	for i:=0;i<len(stack1);i++{
		if stack1[i] != stack2[i] {
			return false
		}
	}

	return true
}
