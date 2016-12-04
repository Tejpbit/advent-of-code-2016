package day4

import (
	"bytes"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type room struct {
	name     string
	number   int
	checksum string
}

type PairList []Pair
type Pair struct {
	Key   rune
	Value int
}

func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int      { return len(p) }
func (p PairList) Less(i, j int) bool {
	if p[i].Value == p[j].Value {
		return p[i].Key > p[j].Key
	}
	return p[i].Value < p[j].Value
}

var reg = regexp.MustCompile("((\\w+-)+)(\\d{3})\\[(\\w*)\\]")

func Run(input string, task int) {
	lines := strings.Split(input, "\n")

	rooms := []room{}
	for _, line := range lines {

		match := reg.FindStringSubmatch(line)
		number, err := strconv.Atoi(match[3])
		if err != nil {
			panic("Bad room id")
		}
		rooms = append(rooms, room{match[1][0 : len(match[1])-1], number, match[4]})

	}

	validRooms := []room{}
	for _, room := range rooms {
		if room.isValidChecksum() {
			validRooms = append(validRooms, room)
		}
	}

	sum := 0
	for _, room := range validRooms {
		sum += room.number
	}

	if task == 1 {
		fmt.Printf("%d", sum)
	} else if task == 2 {
		for _, e := range validRooms {
			fmt.Printf("%d %s\n", e.number, e.decrypt())
		}
	}
}

func (r room) isValidChecksum() bool {
	return r.checksum == r.createChecksum()
}

func (r room) createChecksum() string {
	count := map[rune]int{}
	for _, e := range strings.Replace(r.name, "-", "", -1) {
		c, ok := count[e]
		if ok {
			count[e] = c + 1
		} else {
			count[e] = 1
		}
	}

	s := sortMapByValue(count)
	var buffer bytes.Buffer
	for i := 0; i < 5; i++ {
		buffer.WriteRune(s[i].Key)
	}
	return buffer.String()

}

func sortMapByValue(m map[rune]int) PairList {
	p := make(PairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(p))
	return p
}

func (r room) decrypt() string {
	var buffer bytes.Buffer
	for _, e := range r.name {
		if e == '-' {
			buffer.WriteRune(' ')
		} else {
			buffer.WriteRune(((e - 'a' + rune(r.number)) % rune(26)) + 'a')
		}
	}

	return buffer.String()
}
