package day10

import (
	"fmt"
	"regexp"
	"testing"
)



func Test1(t *testing.T) {

	var match1 = regexp.MustCompile("bot (\\d+) gives low to (output|bot) (\\d+) and high to (output|bot) (\\d+)")
	var match2 = regexp.MustCompile("value (\\d+) goes to bot (\\d+)")

	res1 := match1.FindStringSubmatch("bot 147 gives low to bot 67 and high to bot 71")
	for _, e := range res1 {
		fmt.Printf("%v\n", e)
	}

	res2 := match2.FindStringSubmatch("value 23 goes to bot 76")
	for _, e := range res2 {
		fmt.Printf("%v\n", e)
	}

}
