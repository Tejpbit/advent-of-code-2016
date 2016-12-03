package main

import (
	"flag"
	"fmt"
	"github.com/tejpbit/advent-of-code-2016/day1"
	"github.com/tejpbit/advent-of-code-2016/day2"
	"github.com/tejpbit/advent-of-code-2016/day3"
	"io/ioutil"
)

func main() {
	day := flag.Int("day", 0, "Which challenge to run, default uses latest challenge")
	inputFile := flag.String("input", "", "Which file to take data from and pass to the days challenge")
	task := flag.Int("task", 3, "which part of the daily task to run")
	flag.Parse()


	if *inputFile == "" {
		panic("No input file entered")
	}

	b, err := ioutil.ReadFile(*inputFile)
	stringInput := string(b)
	if err != nil {
		fmt.Printf("%v", err)
		panic("Couldn't read data-day2 file")
	}

	switch *day {
	case 1:
		day1.Run(stringInput, *task)
		return
	case 2:
		day2.Run()
	case 3:
		day3.Run(stringInput, *task)
	}
}
