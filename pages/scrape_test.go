package pages

import (
	"github.com/georgedrummond/learning-go/security"
	"testing"
)

func TestPage(t *testing.T) {
	t.Skip()

	page := Page()

	if page != "hello world" {
		t.Fatalf("Something is broken %x", security.Hash(page))
	}
}
