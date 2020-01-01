package stack

import (
	"strconv"
	"strings"
)

/*
输入:
n = 2
logs =
["0:start:0",
 "1:start:2",
 "1:end:5",
 "0:end:6"]
输出:[3, 4]
说明：
函数 0 在时刻 0 开始，在执行了  2个时间单位结束于时刻 1。
现在函数 0 调用函数 1，函数 1 在时刻 2 开始，执行 4 个时间单位后结束于时刻 5。
函数 0 再次在时刻 6 开始执行，并在时刻 6 结束运行，从而执行了 1 个时间单位。
所以函数 0 总共的执行了 2 +1 =3 个时间单位，函数 1 总共执行了 4 个时间单位

*/
func exclusiveTime(n int, logs []string) []int {
	res := make([]int, n)
	stack := InitMyStack()

	s := strings.Split(logs[0], ":")

	index, err := strconv.Atoi(s[0])
	if err != nil {
		return res
	}
	prev, err := strconv.Atoi(s[2])
	if err != nil {
		return res
	}
	stack.Push(index)

	for i := 1; i < len(logs); i++ {
		arr := strings.Split(logs[i], ":")
		index, _ := strconv.Atoi(arr[0])
		value, _ := strconv.Atoi(arr[2])

		if arr[1] == "start" {
			if !stack.IsEmpty() {
				res[stack.Peek().(int)] += value - prev
			}
			stack.Push(index)
			prev = value
		} else {
			res[stack.Peek().(int)] += value - prev + 1
			stack.Pop()
			prev = value + 1
		}

	}
	return res
}
