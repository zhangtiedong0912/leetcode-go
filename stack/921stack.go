package stack

func minAddToMakeValid(S string) int {
	res := 0
	bal := 0

	for _,v := range S {
		if byte(v) == '('{
			bal ++
		}else{
				bal --
		}
		if bal== -1 {
			res ++
			bal ++
		}
	}
	return res+bal
}
