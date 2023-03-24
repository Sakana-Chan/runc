package system

// import (
// 	"strings"
// )

func FuzzParseStat(data []byte) int {
	_, _ = parseStat(string(data))
	return 1
}