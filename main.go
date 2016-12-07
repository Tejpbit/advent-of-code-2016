package main

import (
	"flag"
	"github.com/tejpbit/advent-of-code-2016/day1"
	"github.com/tejpbit/advent-of-code-2016/day2"
	"github.com/tejpbit/advent-of-code-2016/day3"
	"io/ioutil"
	"log"
	"github.com/tejpbit/advent-of-code-2016/day4"
	"github.com/tejpbit/advent-of-code-2016/day5"
	"github.com/tejpbit/advent-of-code-2016/day6"
	"github.com/tejpbit/advent-of-code-2016/day7"
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
		log.Fatal(err)
	}

	switch *day {
	case 1:
		day1.Run(stringInput, *task)
		return
	case 2:
		day2.Run(stringInput, *task)
	case 3:
		day3.Run(stringInput, *task)
	case 4:
		day4.Run(stringInput, *task)
	case 5:
		day5.Run(stringInput, *task)
	case 6:
		day6.Run(stringInput, *task)
	case 7:
		day7.Run(stringInput, *task)
	}
}
