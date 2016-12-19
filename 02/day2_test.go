package day2

import "testing"

func TestWalkTheLine(t *testing.T) {
	if endCoord := walkTheLine(Coord{1,1}, "ULL", moveOnKeypad); endCoord != (Coord{0,0}) {
		t.Errorf("Bad %v", endCoord)
	}

	if endCoord := walkTheLine(Coord{1,1}, "LURDL", moveOnKeypad); endCoord != (Coord{0,1}) {
		t.Errorf("Bad %v", endCoord)
	}
}

func TestMoveOnKeypad(t *testing.T) {
	if nextCoord := moveOnKeypad(Coord{0,0}, Coord{-1,0}); nextCoord != (Coord{0,0}) {
		t.Errorf("Helo %v", nextCoord)
	}

	if nextCoord := moveOnKeypad(Coord{1,1}, Coord{-1,0}); nextCoord != (Coord{0,1}) {
		t.Errorf("Helo %v", nextCoord)
	}

	if nextCoord := moveOnKeypad(Coord{2,2}, Coord{0,-1}); nextCoord != (Coord{2,1}) {
		t.Errorf("Helo %v", nextCoord)
	}
}


func TestGetDirectionFromChar(t *testing.T) {
	if getDirectionFromChar('L') != (Coord{-1, 0}) {
		t.Error("Bad")
	}

	if getDirectionFromChar('R') != (Coord{1, 0}) {
		t.Error("Bad")
	}

	if getDirectionFromChar('U') != (Coord{0, 1}) {
		t.Error("Bad")
	}

	if getDirectionFromChar('D') != (Coord{0, -1}) {
		t.Error("Bad")
	}
}