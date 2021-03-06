package util

import (
	"fmt"
	"os"
	"strings"
)

func Archive(baseAddr string, dstFolder string, poems []string) {
	for _, v := range poems {
		attributes := strings.Split(v, "-")
		typ := attributes[0]
		school := attributes[1]
		author := attributes[2]
		name := attributes[3][0 : len(attributes[3])-4]
		fmt.Println(typ, school, author, name)
		dstAddr := dstFolder + "/" + typ + "/" + school + "/" + author
		os.MkdirAll(dstAddr, 0766)
		// os.Rename(baseAddr+"/"+v, dstAddr+"/"+v)
		CopyFile(baseAddr+"/"+v, dstAddr+"/"+v)
	}
}
