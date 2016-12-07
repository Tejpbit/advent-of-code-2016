package day5

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func Run(input string, task int) {
	if task == 1 {

		index := 0

		var buffer bytes.Buffer

		for ; buffer.Len() < 8; index += 1 {
			wat := input + strconv.Itoa(index)
			b := md5.Sum([]byte(wat))
			var md5sum string = hex.EncodeToString(b[:])
			if strings.HasPrefix(md5sum, "00000") {
				fmt.Printf("%s\n", md5sum)
				buffer.WriteByte(md5sum[5])
			}
		}

		fmt.Printf("Password is: %s", buffer.String())
	} else if task == 2 {

		var password []rune = make([]rune, 8)
		var found []bool = []bool{false, false, false, false, false, false, false, false}

		for index := 0; !allTrue(found); index += 1 {
			wat := input + strconv.Itoa(index)
			b := md5.Sum([]byte(wat))
			var md5sum string = hex.EncodeToString(b[:])

			fmt.Printf("%b", b)
			if strings.HasPrefix(md5sum, "00000") {
				fmt.Printf("%s\n", md5sum)
				pos, err := strconv.Atoi(string(md5sum[5]))
				if err == nil && pos < 8 && !found[pos] {
					found[pos] = true
					password[pos] = rune(md5sum[6])
					fmt.Printf("%s\n", string(password))
					fmt.Printf("%v\n", found)
				}
			}
		}

		fmt.Printf("Password is: %s", string(password))
	}

}
func allTrue(bools []bool) bool {
	for _, b := range bools {
		if !b {
			return false
		}
	}
	return true
}
