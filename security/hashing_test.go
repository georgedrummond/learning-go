package security

import (
	"testing"
)

func TestHash(t *testing.T) {
	result := Hash("Hello world")
	expected := "1d871aa3aa2af0b598e68618018bbb1d2c4331ab5d31c5df56b148d8a8be3959"

	if result != expected {
		t.Fatalf("Hash is not as expected (expected = %s, got = %s)", result, expected)
	}
}
