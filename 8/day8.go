package day8

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"strconv"
)

type Display [][]bool

func Run(input string, task int) {

	d := createDisplay(50,6)
	for _, line := range strings.Split(input, "\n") {
		err := d.applyInstruction(line)
		fmt.Printf("%s\n", line)
		fmt.Printf("%s\n", d.String())
		if err != nil {
			panic("Bad instruction")
		}
	}
	fmt.Printf("Lit pixel count: %d\n", d.countLitPixels())
	fmt.Printf("%v\n", d.String())
}

func (display Display) countLitPixels() (count int) {

	for _, row := range display {
		for _, pixel := range row {
			if pixel {
				count ++
			}
		}
	}
	return
}

func createDisplay(x, y int) (display Display) {
	display = make([][]bool, y)
	for i, _ := range display {
		display[i] = make([]bool, x)
	}

	return
}

func (display Display) String() string {
	var buffer bytes.Buffer
	for _, row := range display {
		for _, elem := range row {
			if elem {
				buffer.WriteRune('#')
			} else {
				buffer.WriteRune('.')
			}

		}
		buffer.WriteRune('\n')
	}
	return buffer.String()
}



func (display Display) switchRect (x, y int) error {
	if y < 0 || x < 0 || len(display) <= y || len(display[0]) <= x {
		return errors.New("Out of range")
	}

	for i := 0;  i < x; i++ {
		for j := 0; j < y; j++ {
			display[j][i] = !display[j][i]
		}
	}
	return nil
}

func (display Display) rotateColumn (x, by int) error {
	if x < 0 || len(display) == 0 || len(display[0]) <= x {
		return errors.New("Out of range")
	}

	copyColumn := make([]bool, len(display))
	for y := 0; y < len(copyColumn); y++ {
		copyColumn[y] = display[y][x]
	}

	for y := 0; y < len(copyColumn); y++ {
		display[y][x] = copyColumn[(y + (len(copyColumn) - by)) % len(copyColumn) ]
	}

	return nil
}

func (display Display) rotateRow (y, by int) error {
	if y < 0 || len(display) == 0 || y >= len(display) {
		return errors.New("Out of range")
	}

	copyColumn := make([]bool, len(display[0]))
	copy(copyColumn, display[y])
	for x := 0; x < len(copyColumn); x++ {
		display[y][x] = copyColumn[(x + (len(copyColumn) - by)) % len(copyColumn)]
	}
	return nil
}

func (display Display) applyInstruction(instruction string) error {
	strs := strings.Split(instruction, " ")
	if strs[0] == "rect" {
		xy := strings.Split(strs[1], "x")
		x, errX := strconv.Atoi(xy[0])
		y, errY := strconv.Atoi(xy[1])
		if errY != nil || errX != nil{
			fmt.Printf("errY, %v\n", errY)
			fmt.Printf("errX, %v\n", errX)
			panic("bad parsing")
		}
		display.switchRect(x, y)
	} else if strs[0] == "rotate" {
		if strs[1] == "row" {
			y, errY := strconv.Atoi(strs[2][2:])
			by, errBY := strconv.Atoi(strs[4])
			if errY != nil || errBY != nil{
				fmt.Printf("errY, %v\n", errY)
				fmt.Printf("errBY, %v\n", errBY)
				panic("bad parsing")
			}
			display.rotateRow(y, by)
		} else if strs[1] == "column" {
			x, errX := strconv.Atoi(strs[2][2:])
			by, errBY := strconv.Atoi(strs[4])
			if errX != nil || errBY != nil{
				fmt.Printf("errX, %v\n", errX)
				fmt.Printf("errBY, %v\n", errBY)
				panic("bad parsing")
			}
			display.rotateColumn(x, by)
		} else {
			panic("fuck")
		}
	} else {
		return errors.New("Could not handle instruction")
	}
	return nil
}