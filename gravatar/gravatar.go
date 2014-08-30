package gravatar

import (
	"github.com/georgedrummond/learning-go/security"
)

const (
	gravatar_base = "https://gravatar.com/avatar/"
)

func EmailUrl(email string) string {
	hash := security.Hash(email)

	return gravatar_base + hash
}
