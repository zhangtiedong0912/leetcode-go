package stack

/*
"3[a2[bc]]", 返回 "abcbcabcbcabcbc".
*/
type data struct {
	Count int
	Value string
}

func decodeString(s string) string {
	var res string
	var stack []data
	var length int
	var num int

	for _, value := range s {
		if value == '[' {
			//进栈
			stack = append(stack, data{Count: num, Value: res})
			length++
			//初始化
			num = 0
			res = ""
		} else if value == ']' {
			//出栈
			data := stack[length-1]

			//用栈中的数字循环，tmp拼接
			tmp := ""
			for i := data.Count; i > 0; i-- {
				tmp += res
			}
			//拼接res
			res = data.Value + tmp
			//出栈
			stack = stack[:length-1]
			length--
		} else if value >= '0' && value <= '9' {
			//数字
			num = num*10 + int(value-'0')
		} else {
			//字母
			res += string(value)
		}
	}

	return res
}
