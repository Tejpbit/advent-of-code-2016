package day1

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"log"
)

type TurnDirection int

const (
	right = 1
	left  = -1
)

type Compass int

const (
	north = iota
	east  = iota
	south = iota
	west  = iota
)

var currentDirection Compass = north
var x, y = 0, 0
var locationHistory map[Coord]int = make(map[Coord]int)
var firstCoordVisitedTwice Coord
var firstCoordVisitedTwiceSet bool

type Coord struct {
	x int
	y int
}

func Run(input string, task int) {

	instructions := strings.Split(input, ", ")
	for _, instriction := range instructions {
		prefix := instriction[:1]
		if prefix == "R" {
			currentDirection = nextDirection(Compass(currentDirection), right)
		} else if prefix == "L" {
			currentDirection = nextDirection(Compass(currentDirection), left)
		} else {
			panic("Prefix wasn't L or R")
		}

		distance, err := strconv.Atoi(instriction[1:])
		if err != nil {
			panic("fuq")
		}
		addCoords(currentDirection, x, y, distance)
		x, y, err = move(currentDirection, x, y, distance)

	}

	fmt.Printf("Easterbunny HQ location (%d,%d)\n", x, y)
	fmt.Printf("Distance to easterbunny HQ %d\n", abs(x)+abs(y))
	fmt.Printf("First coord visited twice (%d,%d)\n", firstCoordVisitedTwice.x, firstCoordVisitedTwice.y)
	fmt.Printf("Distance to first coord visited twice %d", abs(firstCoordVisitedTwice.x)+abs(firstCoordVisitedTwice.y))
}

func nextDirection(currentDirection Compass, turn TurnDirection) Compass {
	newDirection := Compass((int(currentDirection) + int(turn)) % 4)
	if newDirection < 0 {
		newDirection += 4
	}

	return newDirection
}

func move(direction Compass, xPos int, yPos int, distance int) (int, int, error) {
	if direction == north {
		return xPos, yPos + distance, nil
	} else if direction == east {
		return xPos + distance, yPos, nil
	} else if direction == south {
		return xPos, yPos - distance, nil
	} else if direction == west {
		return xPos - distance, yPos, nil
	} else {
		return xPos, yPos, errors.New("Bad Direciton")
	}
}

func addCoords(direction Compass, xPos int, yPos int, distance int) {
	if distance <= 0 {
		return
	}

	var nextCord Coord = Coord{xPos, yPos}

	if direction == north {
		nextCord = Coord{xPos, yPos + 1}
	} else if direction == east {
		nextCord = Coord{xPos + 1, yPos}
	} else if direction == south {
		nextCord = Coord{xPos, yPos - 1}
	} else if direction == west {
		nextCord = Coord{xPos - 1, yPos}
	} else {
		log.Panicf("Bad direction in addCoords %d", int(direction))
	}

	val, ok := locationHistory[nextCord]
	if ok {
		if !firstCoordVisitedTwiceSet {
			firstCoordVisitedTwiceSet = true
			firstCoordVisitedTwice = nextCord
		}
		val += 1
	} else {
		val = 1
	}
	locationHistory[nextCord] = val
	addCoords(direction, nextCord.x, nextCord.y, distance-1)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
