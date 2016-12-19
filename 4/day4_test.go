package day4

import (
	"testing"
	"fmt"
)

func TestCreateChecksum(t *testing.T) {
	r := room{"not-a-real-room", 404, "oarel"}
	if r.createChecksum() == "oarel" {
		t.Error("Fucked")
	}

}

func TestIsValidChecksum(t *testing.T) {
	r := room{"not-a-real-room", 404, "oarel"}
	fmt.Printf("checksum: %s\n", r.createChecksum())
	if r.isValidChecksum() {
		t.Error("fuck")
	}
}

func TestDecrypt(t *testing.T) {
	r := room{"qzmt-zixmtkozy-ivhz", 343, "zimth"}
	fmt.Printf("%s\n", r.isValidChecksum())
	fmt.Printf("%s\n", r.decrypt())
}