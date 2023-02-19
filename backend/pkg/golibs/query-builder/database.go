package query_builder

import (
	"strconv"
	"strings"
)

// Position argument
func preparePositionalArgsQuery(query string) string {
	parts := strings.Split(query, "?")
	length := len(parts) - 1
	for index := range parts {
		if index < length {
			parts[index] += "$" + strconv.Itoa(index+1)
		}
	}
	return strings.Join(parts, "")
}
