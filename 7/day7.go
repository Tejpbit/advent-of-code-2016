package day7

import (
	"fmt"
	"strings"
)

type IPv7 struct {
	outer []string
	inner []string
	isTLS bool
	isSSL bool
}

func Run(input string, task int) {
	lines := strings.Split(input, "\n")

	if task == 1 {
		countTLS := 0

		for _, line := range lines {
			ipv7, err := parseIPv7(line)
			if err != nil {
				fmt.Errorf("Couldn't parse ipv7 (%s)", line)
			}
			ipv7.determineTLS()
			if ipv7.isTLS {
				countTLS += 1
			}
		}

		fmt.Printf("Count is: %d\n", countTLS)
	} else if task == 2 {
		countSSL := 0

		for _, line := range lines {
			ipv7, err := parseIPv7(line)
			if err != nil {
				fmt.Errorf("Couldn't parse ipv7 (%s)", line)
			}
			ipv7.determineSSL()
			if ipv7.isSSL {
				countSSL += 1
			}

		}
		fmt.Printf("Count is: %d\n", countSSL)

	}
}

func (ipv7 *IPv7) determineTLS() {
	outIsAbba, inIsAbba := false, false
	for _, out := range ipv7.outer {
		if isABBA(out) {
			outIsAbba = true
			break
		}
	}
	for _, in := range ipv7.inner {
		if isABBA(in) {
			inIsAbba = true
			break
		}
	}
	if outIsAbba && !inIsAbba {
		ipv7.isTLS = true
	}
}

func (ipv7 *IPv7) determineSSL() {
	var innerABAs, outerABAs []string
	for _, out := range ipv7.outer {
		ABAs := getABAs(out)
		innerABAs = append(innerABAs, ABAs...)
	}
	for _, in := range ipv7.inner {
		ABAs := getABAs(in)
		outerABAs = append(outerABAs, ABAs...)
	}
	for _, e := range innerABAs {
		var a string = string([]byte{e[1], e[0], e[1]})
		if contains(outerABAs, a) {
			ipv7.isSSL = true
			return
		}
	}
}

func contains(arr []string, e string) bool {
	for _, a := range arr {
		if a == e {
			return true
		}
	}
	return false
}

func isABBA(str string) bool {
	for i, r := range str {
		if i > len(str)-4 {
			break
		} else if r == rune(str[i+3]) && str[i+1] == str[i+2] && str[i] != str[i+1] {
			return true
		}
	}
	return false
}

func getABAs(str string) []string {
	var res []string
	for i, r := range str {
		if i > len(str)-3 {
			break
		} else if r == rune(str[i+2]) {
			res = append(res, str[i:i+3])
		}
	}
	removeDuplicates(&res)
	return res
}

//gotten from https://groups.google.com/d/msg/golang-nuts/-pqkICuokio/ZfSRfU_CdmkJ
func removeDuplicates(xs *[]string) {
	found := make(map[string]bool)
	j := 0
	for i, x := range *xs {
		if !found[x] {
			found[x] = true
			(*xs)[j] = (*xs)[i]
			j++
		}
	}
	*xs = (*xs)[:j]
}

func parseIPv7(ipv7 string) (IPv7, error) {
	var ret IPv7
	ret.outer = []string{}
	ret.inner = []string{}

	for len(ipv7) > 0 {

		outer, inner, rest, err := parseIPv7Part(ipv7)
		if err != nil {
			return ret, err
		}
		ret.outer = append(ret.outer, outer)
		if len(rest) > 0 {
			ret.inner = append(ret.inner, inner)

		}
		ipv7 = rest
	}
	return ret, nil
}

func parseIPv7Part(ipv7Part string) (string, string, string, error) {
	i1 := strings.IndexRune(ipv7Part, '[')
	i2 := strings.IndexRune(ipv7Part, ']')
	if i1 == -1 && i2 != -1 || i1 != -1 && i2 == -1 {
		return "", "", "", fmt.Errorf("Bad ipv7 string: %s", ipv7Part)
	} else if i1 == -1 && i2 == -1 {
		return ipv7Part, "", "", nil
	}

	outsideBrackets := ipv7Part[:i1]
	insideBrackets := ipv7Part[i1+1 : i2]
	remainder := ipv7Part[i2+1:]
	return outsideBrackets, insideBrackets, remainder, nil
}
