package day9

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/Workiva/go-datastructures/queue"
	"time"
)

type ExpanderElement struct {
	length   int
	repeat   int
	position int // position after the expander element eg XXX(4x3)XXX
} //                                                ^

func Run(input string, task int) {
	/*inputQueue := queue.New(10)
	inputChannel := make(chan rune, len(input))
	ready := make(chan bool)
	go func() {
		for _, r := range input {
			inputChannel <- r
		}
		close(inputChannel)
		ready <- true
	}()

	for _, r := range input {
		inputQueue.Put(r)
	}

	println("Not Ready")
	<-ready
	fmt.Println("GoGoGo")
	outputChannel := expand(inputChannel)

	count := 0
	for range outputChannel {
		count++
	}
	fmt.Printf("%d\n", count)*/
}

func expand(in *queue.Queue) []rune {
	out := make([]rune, 1000000000)
	localQueue := queue.New(1)

	currentWork := []rune{}

	// function for reading next rune
	// prioritise localQueue before in
	readNext := func(internal, external *queue.Queue) (rune, bool) {
		r, err := internal.Poll(1, time.Nanosecond)
		if err == nil {
			if ru, ok := r[0].(rune); ok {
				return ru, true
			} else {
				panic("Not a rune1?")
			}
		} else {
			r, err = external.Poll(1, time.Nanosecond)
			if err == nil {
				if ru, ok := r[0].(rune); ok {
					return ru, true
				} else {
					panic("Not a rune2?")
				}
			} else {
				return 'a', false
			}
		}
		return 'a', false
	}

	readUntil := func(prio, norm *queue.Queue, checker func(r rune) bool) []rune {
		acc := []rune{}

		for {
			r, ok := readNext(prio, norm)
			if !ok {
				break
			}
			acc = append(acc, r)
			if checker(r) {
				break
			}
		}
		return acc
	}

	for {
		nextRune, ok := readNext(localQueue, in)
		if !ok {
			break
		}
		currentWork = []rune{}
		if nextRune == '(' {
			currentWork = append(currentWork, '(')
			currentWork = append(currentWork, readUntil(localQueue, in, func(r rune) bool {return (r == ')')})...)
			stringData := string(currentWork[1 : len(currentWork)-1])
			length, repeat, _ := parseCompressionMetadata(stringData)
			currentWork = []rune{}
			// read 'length' number of runes from input
			for i := 0; i < length; i++ {
				n, _ := readNext(localQueue, in)
				currentWork = append(currentWork, n)
			}

			toWriteToInternal := []rune{}
			for i := 0; i < repeat; i++ {
				toWriteToInternal = append(toWriteToInternal, currentWork...)
			}

			for i := 0; i < len(toWriteToInternal); i++ {
				localQueue.Put(toWriteToInternal[i])
			}

		} else {
			if len(out) % 100000 == 0 {
				fmt.Printf("len: %d\tcap: %d\n", len(out), cap(out))
			}
			out = append(out, nextRune)
		}
	}

	return out
}

func readUntil(c chan rune, expected rune) (acc []rune) {
	acc = []rune{}
	for r := range c {
		acc = append(acc, r)
		if r == expected {
			return
		}
	}
	return
}

func split(input string) []string {
	var split []string
	/*
		nuvarandeIndex
		om nuvarande index inte börjar med (,
			hitta nästa (,
				finns det inget (?
				appenda från nuvarande index till slutet till spliten
			ta substrungen fram tills (
			appenda till split
			uppdatera nuvarande index till positionen för (
		om nuvarande index börjar med (,
			hitta nästa )
			parsa ut length
			ta substringen från ( till ) +1 + length
			appenda till split

	*/
	currentIndex := 0
	for currentIndex < len(input) {
		if input[currentIndex] != '(' {
			left := findInString(input, '(', currentIndex)
			if left == -1 {
				split = append(split, input[currentIndex:])
				break
			}
			split = append(split, input[currentIndex:left])
			currentIndex = left
		} else {
			left := currentIndex
			right := findInString(input, ')', currentIndex)
			length, _, _ := parseCompressionMetadata(input[left+1 : right])
			split = append(split, input[left:right+1+length])
			currentIndex = right + 1 + length
		}

	}
	return split
}

func extractExpanderElements(input string) (expanders []ExpanderElement) {
	for i := 0; i < len(input); i++ {
		r := input[i]
		if r == '(' {
			leftBracketIndex := i
			rightBracketIndex := findInString(input, ')', i)
			length, repeat, err := parseCompressionMetadata(input[leftBracketIndex+1 : rightBracketIndex])
			if err != nil {
				panic("bad parse on compression metadata")
			}
			expanders = append(expanders, ExpanderElement{length, repeat, rightBracketIndex + 1})
		}
	}
	return
}

func expandFirst(input string) (str string) {
	left := strings.IndexRune(input, '(')
	right := strings.IndexRune(input, ')')
	if right == -1 {
		return input
	}
	length, repeat, err := parseCompressionMetadata(input[left+1 : right])
	if err != nil {
		panic("bad parsing from compression metadata")
	}
	head := input[:left]
	middle := strings.Repeat(input[right+1:right+1+length], repeat)
	tail := input[right+1+length:]
	return head + middle + tail

}

func expandLast(input string) (string, error) {
	right := lastIndex(input, ")")
	left := lastIndex(input, "(")
	if right < 0 || left < 0 {
		return "", errors.New(fmt.Sprintf("left and right is %d, %d", left, right))
	}
	if right <= left {
		return "", errors.New("Right is left of left")
	}
	length, repeat, err := parseCompressionMetadata(input[left+1 : right])
	if err != nil {
		panic("bad parsing from compression metadata")
	}
	newTail := strings.Repeat(input[right+1:right+1+length], repeat)

	return input[:left] + newTail + input[right+1+length:], nil
}

func parseCompressionMetadata(stringData string) (length, repeat int, err error) {
	i := strings.IndexRune(stringData, 'x')
	if i == -1 {
		return 0, 0, errors.New("could not parse stringData")
	}
	length, err = strconv.Atoi(stringData[:i])
	if err != nil {
		return 0, 0, errors.New("could not parse stringData")
	}
	repeat, err = strconv.Atoi(stringData[i+1:])
	if err != nil {
		return 0, 0, errors.New("could not parse stringData")
	}
	return
}

func findInString(str string, r rune, startIndex int) int {
	if startIndex >= len(str) {
		return -1
	}
	for i := startIndex; i < len(str); i++ {
		if rune(str[i]) == r {
			return i
		}
	}
	return -1
}

func lastIndex(s string, sub string) int {
	for i := len(s) - len(sub); i >= 0; i-- {
		if s[i:i+len(sub)] == sub {
			return i
		}
	}
	return -1
}

func countExpanded(input string) int {
	sum := 0
	for {
		if input == "" {
			return sum
		}
		count, next := countAndRemove(input)
		sum += count
		input = next
	}
}

func countAndRemove(input string) (count int, rest string) {
	left := strings.IndexRune(input, '(')
	right := strings.IndexRune(input, ')')
	if left == -1 || right == -1 {
		return len(input), ""
	}
	length, repeat, err := parseCompressionMetadata(input[left+1 : right])
	if err != nil {
		return len(input), ""

	}
	return left + length*repeat, input[right+1+length:]

}
