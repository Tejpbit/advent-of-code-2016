package day1

import (
	"testing"
)

func TestTurning(t *testing.T) {
	dir := nextDirection(north, right)
	if dir != east {
		t.Errorf("Expected %d got %d", east, dir)
	}

	dir = nextDirection(north, left)
	if dir != west {
		t.Errorf("Expected %d got %d", west, dir)
	}

	dir = nextDirection(west, right)
	if dir != north {
		t.Errorf("Expected %d got %d", north, dir)
	}

}

func TestMove(t *testing.T) {
	x, y := 0, 0
	x, y, _ = move(north, x, y, 4)
	if !(x == 0 && y == 4) {
		t.Errorf("Expected %d,%d got %d,%d", 0, 4, x, y)
	}
	x, y, _ = move(west, x, y, 4)
	if !(x == -4 && y == 4) {
		t.Errorf("Expected %d,%d got %d,%d", -4, 4, x, y)
	}
	x, y, _ = move(south, x, y, 9)
	if !(x == -4 && y == -5) {
		t.Errorf("Expected %d,%d got %d,%d", -4, -5, x, y)
	}
	x, y, _ = move(east, x, y, 2)
	if !(x == -2 && y == -5) {
		t.Errorf("Expected %d,%d got %d,%d", -2, -5, x, y)
	}
}

func TestAddCoord(t *testing.T) {
	if len(locationHistory) != 0 {
		t.Error("Location history isn't empty")
	}
	x, y := 0, 0
	addCoords(north, x, y, 10)
	if len(locationHistory) != 10 {
		t.Error("y tho?")
	}

}
