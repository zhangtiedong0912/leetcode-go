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
	n :=removeKdigits("1432219", 3)
	fmt.Println(n)
}
