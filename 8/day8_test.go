package day8

import (
	"fmt"
	"testing"
	"strings"
)

func TestCreateAndDisplay(t *testing.T) {
	d := createDisplay(60, 5)
	d.switchRect(7, 4)
	pixelCount := d.countLitPixels()
	if pixelCount != 28 {
		t.Error("Not 28 :(")
	}
	d.rotateRow(2, 5)
	d.rotateColumn(6, 5)
	if d.countLitPixels() != pixelCount {
		t.Error("Should be same")
	}
	fmt.Printf("%s\n", d.String())
	d.rotateRow(0, 1)
	fmt.Printf("%s\n", d.String())

}

func TestRotateColumn(t *testing.T) {
	d := createDisplay(60, 5)
	d.switchRect(7, 3)
	d.rotateColumn(5, 2)
	fmt.Printf("%s\n\n", d.String())
}

func TestRotateRow(t *testing.T) {
	d := createDisplay(60, 5)
	d.switchRect(7, 3)
	d.rotateRow(1, 4)
	d.switchRect(5, 4)
	fmt.Printf("%s\n", d.String())
}

func TestRectInstruction(t *testing.T) {
	d := createDisplay(60, 5)
	d.applyInstruction("rect 4x2")
	d.applyInstruction("rotate row y=0 by 5")
	d.applyInstruction("rotate column x=0 by 1")
	fmt.Printf("%v\n", d.String())
}

func TestExample(t *testing.T) {
	d := createDisplay(7, 3)

	input := `rect 3x2
rotate column x=1 by 1
rotate row y=0 by 4
rotate column x=1 by 1
rotate column x=0 by 1
rotate row x=0 by 1`

	lines := strings.Split(input, "\n")

	for _, i := range lines {
		d.applyInstruction(i)
	}
	fmt.Printf("%s\n", d.String())
	fmt.Printf("%d\n", d.countLitPixels())

}