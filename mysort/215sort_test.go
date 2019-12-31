package mysort

import (
	"fmt"
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	nums := []int{3, 2, 1, 5, 6, 4}
	n := findKthLargest(nums, 2)
	fmt.Println(n)
}
