package stack

/*
输入：S = "leet2code3", K = 10
输出："o"
*/
func decodeAtIndex(S string, K int) string {
	num := 0
	for _, v := range S {
		if byte(v) >= '2' && byte(v) <= '9' {
			num *= int(byte(v) - '0')
		} else {
			num++
		}
	}
	if num < K {
		return ""
	}

	l := len(S)
	for i:=l-1;i>=0;i--{
		K %=num
		if K==0 && (byte(S[i]) < '2' || byte(S[i]) > '9') {
			return string(byte(S[i]))
		}
		if byte(S[i]) >= '2' && byte(S[i]) <= '9' {
			num /= int(byte(S[i]) - '0')
		}else{
			num --
		}
	}

	return ""
}


