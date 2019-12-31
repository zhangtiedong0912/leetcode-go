package stack

import "strings"

/*
思路：把路径按照 `/` 切割，遍历string数组,遇到`.`和`` 不放入，遇到`..` 出栈，遇到其他放入栈
注意出栈的时候要判断栈是否为空，若为空跳过本次循环（到了根目录）
*/
func simplifyPath(path string) string {
	var res = ""
	if len(path) == 0 {
		return `/`
	}

	var stack []string
	length := 0

	array := strings.Split(path, `/`)
	for _, value := range array {
		if value == "." || value == "" {
			continue
		}
		if value == ".." {
			if length == 0 {
				continue
			}
			stack = stack[:length-1]
			length--
		} else {
			stack = append(stack, value)
			length++
		}
	}

	if len(stack) == 0 {
		return `/`
	}
	for _, value := range stack {
		res += `/` + value
	}

	return res
}
