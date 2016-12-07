package day5

import (
	"testing"
)

func TestAllTrue(t *testing.T) {
	if allTrue([]bool{false,false,false}) {
		t.Error("Shouldn't be true")
	}

	if allTrue([]bool{true,false,false}) {
		t.Error("Shouldn't be true")
	}

	if !allTrue([]bool{true,true,true,true}) {
		t.Error("Shouldn't be true")
	}

	if allTrue([]bool{true,true,true,false}) {
		t.Error("Shouldn't be true")
	}

	if !allTrue([]bool{}) {
		t.Error("Shouldn't be true")
	}

	if allTrue([]bool{true, true, true, true, true, true, false, true}) {
		t.Error("Shouldn't be true")
	}


}

func BenchmarkTask1(b *testing.B) {
	b.StartTimer()
	Run("ojvtpuvg", 1)
	b.StopTimer()

}

func BenchmarkTask2(b *testing.B) {
	b.StartTimer()
	Run("ojvtpuvg", 2)
	b.StopTimer()

}