package day3

import (
	"fmt"
	"testing"
)

func TestZip3(t *testing.T) {
	expected := [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}
	res := zip3([]int{1, 2, 3, 8}, []int{4, 5, 6}, []int{7, 8, 9})

	fmt.Printf("%v\n", expected)
	fmt.Printf("%v", res)
}

func TestParse2(t *testing.T) {
	testString := `101 301 501
	102 302 502
	103 303 503
	201 401 601
	202 402 602
	203 403 603`

	fmt.Printf("%v", parseInput2(testString))
}
