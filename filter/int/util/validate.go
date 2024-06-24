package util

import (
	"strings"
)

func ValidateString(input string) (validate bool) {

	if strings.TrimSpace(input) == "" {
		return false

	}

	return true

}
