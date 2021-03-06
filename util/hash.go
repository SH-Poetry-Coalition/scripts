package util

import (
	"hash/fnv"
	"strconv"
)

func Hash(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	return strconv.Itoa(int(h.Sum32()))
}
