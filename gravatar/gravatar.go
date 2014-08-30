package gravatar

import (
	"./../security"
)

func EmailUrl(email string) string {
	hash := security.Hash(email)

	return "https://gravatar.com/avatar/" + hash
}
