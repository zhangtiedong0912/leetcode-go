package stack

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	ss := "1234"
	fmt.Println(ss[0:4])
	n := isValid("()[]{}")
	fmt.Println(n)
}

func TestSimplifyPath(t *testing.T) {
	n := simplifyPath(`/a/../../b/../c//.//`)
	fmt.Println(n)
}

func TestDecodeString(t *testing.T) {
	fmt.Println(decodeString("3[a2[c]]"))
}

func TestRemoveKdigits(t *testing.T) {
	n := removeKdigits("1432219", 3)
	fmt.Println(n)
}

func TestNextGreaterElements(t *testing.T) {
	n := nextGreaterElements([]int{1, 2, 1})
	fmt.Println(n)
}

func TestExclusiveTime(t *testing.T) {

	input := []string{"0:start:0",
		"1:start:2",
		"1:end:5",
		"0:end:6"}
	res := exclusiveTime(2, input)
	fmt.Println(res)
}

func TestDailyTemperatures(t *testing.T) {
	res := dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})
	fmt.Println(res)
}

func TestDecodeAtIndex(t *testing.T) {
	res := decodeAtIndex("leet2code3", 6)
	fmt.Println(res)
	res = decodeAtIndex("leet2code3", 10)
	fmt.Println(res)
	res = decodeAtIndex("ha22", 5)
	fmt.Println(res)
	res = decodeAtIndex("aw4eguc6cs", 41)
	fmt.Println(res)
}

func TestSumSubarrayMins(t *testing.T) {
	res := sumSubarrayMins([]int{3,1,2,4})
	fmt.Println(res)
}

