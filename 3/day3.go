package day3

import (
	"fmt"
	"strconv"
	"strings"
)

func Run(input string, task int) {
	if task == 1 {
		triangles := parseInput1(input)
		count := countPossibleTriangles(triangles)
		fmt.Printf("Task1, triangle count %d", count)

	} else if task == 2 {
		triangles := parseInput2(input)
		res := countPossibleTriangles(triangles)
		fmt.Printf("Task2, possible triangle count %d", res)
	}

}

func countPossibleTriangles(triangles [][]int) (count int) {
	for _, t := range triangles {
		if t[0]+t[1] > t[2] && t[0]+t[2] > t[1] && t[1]+t[2] > t[0] {
			count++
		}
	}
	return
}

func parseInput1(s string) [][]int {
	lines := strings.Split(s, "\n")

	triangles := make([][]int, len(lines))
	for i, line := range lines {
		triangles[i] = toIntArray(strings.Fields(line))
	}
	return triangles
}

func parseInput2(s string) [][]int {
	lines := strings.Split(s, "\n")
	intLines := make([][]int, len(lines))
	for i, line := range lines {
		intLines[i] = toIntArray(strings.Fields(line))
	}

	result := [][]int{}
	for i := 0; i < len(lines); i += 3 {
		z := zip3(intLines[i], intLines[i+1], intLines[i+2])
		result = append(result, z...)
	}
	return result
}

func zip3(arr1 []int, arr2 []int, arr3 []int) (result [][]int) {
	min := min3(len(arr1), len(arr2), len(arr3))
	result = make([][]int, min)
	for i := 0; i < min; i++ {
		nextZip := make([]int, 3)
		nextZip[0] = arr1[i]
		nextZip[1] = arr2[i]
		nextZip[2] = arr3[i]
		result[i] = nextZip
	}
	return
}

func min3(a int, b int, c int) (min int) {
	min = a
	if b < a {
		min = b
	}
	if c < b {
		min = c
	}
	return
}

func toIntArray(sArr []string) []int {
	res := []int{}
	for _, e := range sArr {
		i, err := strconv.Atoi(e)
		if err != nil {
			fmt.Printf("\n%v\n", sArr)
			panic("Couldn't convert string array to int array")
		}
		res = append(res, i)
	}
	return res
}
