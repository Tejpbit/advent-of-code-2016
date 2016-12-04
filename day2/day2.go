package day2

import (
	"fmt"
	"strings"
)

type Direction Coord

type Coord struct {
	x int
	y int
}

var left = Coord{-1, 0}
var right = Coord{1, 0}
var up = Coord{0, -1}
var down = Coord{0, 1}

var currentCoord Coord = Coord{0, 2}
var keypad1 [][]int = [][]int{
	{1, 2, 3},
	{4, 5, 6},
	{7, 8, 9}}

var keypad2 [][]int = [][]int{
	{-0x1, -0x1, 0x1, -0x1, -0x1},
	{-0x1, 0x2, 0x3, 0x4, -0x1},
	{0x5, 0x6, 0x7, 0x8, 0x9},
	{-0x1, 0xA, 0xB, 0xC, -0x1},
	{-0x1, -0x1, 0xD, -0x1, -0x1},
}
var combination []int

func Run(input string, task int) {
	inputLines := strings.Split(input, "\n")

	keypad := keypad1
	keypadMove := moveOnKeypad
	if task == 2 {
		keypad = keypad2
		keypadMove = moveOnKeypad2
	}

	for _, line := range inputLines {
		resultingCoord := walkTheLine(currentCoord, line, keypadMove)
		currentCoord = resultingCoord
		combination = append(combination, keypad[resultingCoord.y][resultingCoord.x])
	}
	for _, e := range combination {
		fmt.Printf("%X", e)
	}
}

func walkTheLine(startCoord Coord, line string, moveFunction func(c Coord, m Coord) (Coord)) (currentCoord Coord) {
	currentCoord = startCoord
	for _, char := range line {
		direction := getDirectionFromChar(char)
		currentCoord = moveFunction(currentCoord, direction)
	}
	return
}

func moveOnKeypad(currentPos Coord, move Coord) Coord {
	cord := Coord{currentPos.x + move.x, currentPos.y + move.y}

	if cord.x < 0 {
		cord.x = 0
	} else if cord.x > 2 {
		cord.x = 2
	}

	if cord.y < 0 {
		cord.y = 0
	} else if cord.y > 2 {
		cord.y = 2
	}
	return cord
}

func moveOnKeypad2(currentPos Coord, move Coord) (cord Coord) {
	cord = Coord{currentPos.x + move.x, currentPos.y + move.y}

	if cord.x < 0 {
		cord.x = 0
	} else if cord.x > 4 {
		cord.x = 4
	}

	if cord.y < 0 {
		cord.y = 0
	} else if cord.y > 4 {
		cord.y = 4
	}

	if keypad2[cord.y][cord.x] == -0x1 {
		cord = currentPos
	}

	return cord
}

func getDirectionFromChar(c rune) Coord {
	return map[rune]Coord{
		'L': left,
		'R': right,
		'U': up,
		'D': down,
	}[c]
}
