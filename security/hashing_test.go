package security

import (
	"testing"
)

func TestHash(t *testing.T) {
	result := Hash("Hello world")
	expected := "3e25960a79dbc69b674cd4ec67a72c62"

	if result != expected {
		t.Fatalf("Hash is not as expected (expected = %s, got = %s)", result, expected)
	}
}
