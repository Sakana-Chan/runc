package system

import (
	"strings"
)

func FuzzParseStat(data []byte) int {
	_, _ := parseStat(strings.NewReader(string(data)))
	return 1
}