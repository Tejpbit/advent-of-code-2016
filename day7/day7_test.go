package day7

import (
	"testing"
	"fmt"
)

func TestParseIPv7(t *testing.T) {


	a1, _ := parseIPv7("gmmfbtpprishiujnpdi[wedykxqyntvrkfdzom]uidgvubnregvorgnhm")
	a2, _ := parseIPv7("txxplravpgztjqcw[txgmmtlhmqpmmwp]bmhfgpmafxqwtrpr[inntmjmgqothdzfqgxq]cvtwvembpvdmcvk")
	a3, _ := parseIPv7("gkxjhpayoyrrpcr[mwyoahlkqyhtznyzrm]mvmurvsrgjunjjepn[mkoumuohilpcfgbmsmh]hpwggyvjkusjxcyojyr[wqxyuzbewpjzlyqmkhw]nniczueulxtdsmkniex")
	a4, _ := parseIPv7("vuzyoofrvaanszwndyt[mzcbhmabgnetrpje]tqnygwhmwrbyosbke[gehqzyhlnyufknqmueo]ngendggbjcvazwol")

	fmt.Printf("ips %v\n%v\n%v\n%v\n", a1, a2, a3, a4)

}


func TestHasAbba(t *testing.T) {

	ret := isABBA("gmmfbtpptishiujnpdi")
	if !ret {
		t.Error("Bad")
	}
	ret = isABBA("txxtlravpgztjqcw")
	if !ret {
		t.Error("Bad")
	}
	ret = isABBA("vuzyoofrvaanszwndtyyt")
	if !ret {
		t.Error("Bad")
	}
}

func TestDetermineTSL(t *testing.T) {
	a1, _ := parseIPv7("gmmfbtpptishiujnpdi[wedykxqyntvrkfdzom]uidgvubnregvorgnhm")
	fmt.Printf("before: %v\n", a1)
	a1.determineTLS()
	fmt.Printf("after : %v\n", a1)

	fmt.Printf("%v", a1)
	if !a1.isTLS {
		t.Error("Should be TSL")
	}

}

func TestGetABAs(t *testing.T) {
	fmt.Printf("%v", getABAs("gmmfbtptptishsujnpdi"))

}

func TestDetermineSSL(t *testing.T) {
	a2, _ := parseIPv7("aba[bab]xyz")
	a2.determineSSL()
	fmt.Printf("%v", a2)
	if !a2.isSSL {
		t.Error("Should be SSL")
	}
}