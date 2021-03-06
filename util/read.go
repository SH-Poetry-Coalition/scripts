package util

import (
	"os"
	"path"
)

func ReadPoems(baseAddr string) (poems []string) {
	files, _ := os.ReadDir(baseAddr)
	for _, v := range files {
		if path.Ext(v.Name()) == ".txt" {
			poems = append(poems, v.Name())
		}
	}
	return poems
}
