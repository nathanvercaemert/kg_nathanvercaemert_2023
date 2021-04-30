package main

import "testing"

var testArray []string = []string{"123", "0", "987", "456"}
var expectedOutput string = "OneTwoThree, Zero, NineEightSeven, FourFiveSix"

func TestPhoneticCurrent(t *testing.T) {

	got := PhoneticCurrent(testArray, 4)
	want := expectedOutput

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func BenchmarkPhoneticCurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PhoneticCurrent(testArray, 4)
	}
}

func TestPhoneticPotentialImprovement(t *testing.T) {

	got := PhoneticPotentialImprovement(testArray, 4)
	want := expectedOutput

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func BenchmarkPhoneticPotentialImprovement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PhoneticPotentialImprovement(testArray, 4)
	}
}
