package main

import "testing"

var testArray []string = []string{"123", "0", "987", "456"}
var expectedOutput string = "OneTwoThree, Zero, NineEightSeven, FourFiveSix"

// Tests the correctness of current implementation.
func TestPhoneticCurrent(t *testing.T) {

	got := PhoneticCurrent(testArray, 4)
	want := expectedOutput

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

// Tests the efficiency of current implementation.
func BenchmarkPhoneticCurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PhoneticCurrent(testArray, 4)
	}
}

// Tests the correctness of potential improvement.
func TestPhoneticPotentialImprovement(t *testing.T) {

	got := PhoneticPotentialImprovement(testArray, 4)
	want := expectedOutput

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

// Tests the efficiency of potential improvement.
func BenchmarkPhoneticPotentialImprovement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PhoneticPotentialImprovement(testArray, 4)
	}
}
