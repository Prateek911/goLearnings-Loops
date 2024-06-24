package util

import (
	"strings"
)

func ValidateString(input string) (validate bool) {

	return strings.TrimSpace(input) == ""

}
