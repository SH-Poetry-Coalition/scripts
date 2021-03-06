package util

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func ReadCorrespondence(reverseCorrespondenceFile string) map[string]string {
	reverseCorrespondence := make(map[string]string)
	fin, err := os.Open(reverseCorrespondenceFile)
	if err != nil {
		panic(err)
	}
	br := bufio.NewReader(fin)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		lineList := strings.Split(string(line), " ")
		// fmt.Println(lineList[0], lineList[1])
		reverseCorrespondence[lineList[0]] = lineList[1]
	}
	return reverseCorrespondence
}

func Reverse(baseAddr string, dstAddr string, reverseCorrespondence map[string]string) {
	for i, v := range reverseCorrespondence {
		os.MkdirAll(dstAddr, 0766)
		CopyFile(baseAddr+"/"+i, dstAddr+"/"+v)
	}
}
