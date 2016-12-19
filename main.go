package main

import (
	"flag"
	"github.com/tejpbit/advent-of-code-2016/01"
	"github.com/tejpbit/advent-of-code-2016/02"
	"github.com/tejpbit/advent-of-code-2016/3"
	"io/ioutil"
	"log"
	"github.com/tejpbit/advent-of-code-2016/4"
	"github.com/tejpbit/advent-of-code-2016/5"
	"github.com/tejpbit/advent-of-code-2016/6"
	"github.com/tejpbit/advent-of-code-2016/7"
	"github.com/tejpbit/advent-of-code-2016/8"
	"github.com/tejpbit/advent-of-code-2016/9"
	"github.com/tejpbit/advent-of-code-2016/10"
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
	case 8:
		day8.Run(stringInput, *task)
	case 9:
		day9.Run(stringInput, *task)
		//task 2 453538 is to low
	case 10:
		day10.Run(stringInput, *task)
	}
}
