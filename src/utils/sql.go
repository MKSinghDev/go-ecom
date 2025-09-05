package utils

import (
	"fmt"
	"strings"
)

func BuildPostgreSQLPlaceholders(ids []int) string {
	if len(ids) == 0 {
		return ""
	}

	placeholders := make([]string, len(ids))
	for i := range ids {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
	}
	return strings.Join(placeholders, ",")
}
