package stack

func isValid(s string) bool {
	symbol := map[byte]byte{'}': '{', ']': '[', ')': '('}
	stack := []int32{}
	length := 0

	for _, v := range s {
		if value, ok := symbol[byte(v)]; ok {
			if length == 0 {
				return false
			}
			if value != byte(stack[length-1]) {
				return false
			}
			stack = stack[0 : length-1]
			length--
		} else {
			stack = append(stack, v)
			length++
		}
	}

	return length == 0
}
