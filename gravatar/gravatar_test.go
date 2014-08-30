package gravatar

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestEmailUrl(t *testing.T) {
	gravatar_url := EmailUrl("georgedrummond@gmail.com")
	expected := "https://gravatar.com/avatar/d278a12b969a495ab16fdd942e748fe5"

	assert.Equal(t, gravatar_url, expected)
}
