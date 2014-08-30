package gravatar

import (
	"github.com/georgedrummond/learning-go/security"
	"strconv"
)

const (
	gravatar_base = "https://gravatar.com/avatar/"
)

type Profile struct {
	Email string
	Size  int
}

func (p *Profile) Url() string {
	var size_param string

	if p.Size != 0 {
		size_param = "?s=" + strconv.Itoa(p.Size)
	}

	return gravatar_base + security.Hash(p.Email) + size_param
}
