package uuid32

import (
	"strings"

	"github.com/google/uuid"
)

func New() string {
	return strings.ReplaceAll(uuid.NewString(), "-", "")
}

func S() string {
	return New()
}
