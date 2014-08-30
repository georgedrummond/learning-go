package gravatar

import (
	"testing"
)

func TestEmailUrl(t *testing.T) {
	gravatar_url := EmailUrl("georgedrummond@gmail.com")
	expected := "https://gravatar.com/avatar/d278a12b969a495ab16fdd942e748fe5"

	if gravatar_url != expected {
		t.Fatalf("Incorrect gravatar url formed (actual=%s, expected=%s)", gravatar_url, expected)
	}
}
