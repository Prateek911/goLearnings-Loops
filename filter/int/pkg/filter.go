package pkg

import "strings"

func Filter(content []byte, searchString string) int {
	count := strings.Count(string(content), searchString)

	return count
}
