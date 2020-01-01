package stack

func asteroidCollision(asteroids []int) []int {
	stack := []int{}

	for i:=0;i<len(asteroids);i++{
		a := asteroids[i]
		if a > 0 {
			stack = append(stack, a)
		}else if len(stack) == 0 || stack[len(stack)-1] < 0 {
			stack = append(stack, a)
		}else if -a > stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			i--
		}else if -a == stack[len(stack)-1]{
			stack = stack[:len(stack)-1]
		}
	}

	return stack
}
