package syntax

import (
	"strings"
)

func CStringStrip(s string) string {
	if i := strings.Index(s, "\x00"); i >= 0 {
		return s[:i]
	}

	return s
}
