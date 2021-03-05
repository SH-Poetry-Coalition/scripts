package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func main() {
	baseAddr := "."
	var poems []string
	files, _ := os.ReadDir(baseAddr)
	for _, v := range files {
		if path.Ext(v.Name()) == ".txt" {
			poems = append(poems, v.Name())
		}
	}
	fmt.Println(poems)

	for _, v := range poems {
		attributes := strings.Split(v, "-")
		typ := attributes[0]
		school := attributes[1]
		author := attributes[2]
		name := attributes[3][0 : len(attributes[3])-4]
		fmt.Println(typ, school, author, name)
		dstAddr := baseAddr + "/" + typ + "/" + school + "/" + author
		os.MkdirAll(dstAddr, 0766)
		os.Rename(baseAddr+"/"+v, dstAddr+"/"+v)
	}
}
