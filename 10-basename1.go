
/*
func basename(s string) string {
	// discard prefix
	for i := len(s)-1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	// discard suffix
	for i := len(s)-1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
		}
	}
	return s
}
*/
//  |
//  |
//  V

package main
import (
	"strings"
)

func basename(s string) string {
	slash = strings.LastIndex(s, "/")  // -1 if "/" not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}

	return s
}
