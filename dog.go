// Package dog transforms the given human years to dog years.
package dog

import (
	"strings"
)

func WhenGrowUp(s string) string {
	return "When the puppy grows up it says: " + strings.ToUpper(s)
}

func Years(i int) int {
	return i * 7
}
