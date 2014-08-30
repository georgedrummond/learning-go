package gravatar

import (
	"github.com/georgedrummond/learning-go/security"
)

func EmailUrl(email string) string {
	hash := security.Hash(email)

	return "https://gravatar.com/avatar/" + hash
}
