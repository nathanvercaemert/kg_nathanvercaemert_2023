package main

import "testing"

var testArray []int = []int{123, 0, 987456}
var expectedOutput string = "OneTwoThree, Zero, NineEightSevenFourFiveSix"

func TestPhoneticCurrent(t *testing.T) {

	got := PhoneticCurrent(testArray, 3)
	want := expectedOutput

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func BenchmarkPhoneticCurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PhoneticCurrent(testArray, 3)
	}
}

func TestPhoneticPotentialImprovement(t *testing.T) {

	got := PhoneticPotentialImprovement(testArray, 3)
	want := expectedOutput

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func BenchmarkPhoneticPotentialImprovement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PhoneticPotentialImprovement(testArray, 3)
	}
}
