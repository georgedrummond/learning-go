package gravatar

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestUrl(t *testing.T) {
	expected := "https://gravatar.com/avatar/d278a12b969a495ab16fdd942e748fe5"
	profile := Profile{Email: "georgedrummond@gmail.com"}

	assert.Equal(t, profile.Size, 0)

	assert.Equal(t, profile.Url(), expected)
}

func TestUrlWithSize(t *testing.T) {
	expected := "https://gravatar.com/avatar/d278a12b969a495ab16fdd942e748fe5?s=300"
	profile := Profile{Email: "georgedrummond@gmail.com", Size: 300}

	assert.Equal(t, profile.Url(), expected)
}
