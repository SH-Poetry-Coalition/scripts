package util

import (
	"os"
)

func GenCorrespondence(baseAddr string, poems []string) (map[string]string, map[string]string) {
	// baseAddr := "./poems_original"
	correspondence := make(map[string]string)
	reverseCorrespondence := make(map[string]string)
	for _, v := range poems {
		anonymousName := Hash(v) + ".txt"
		correspondence[v] = anonymousName
		reverseCorrespondence[anonymousName] = v
	}
	return correspondence, reverseCorrespondence
}

func WriteCorrespondence(correspondence map[string]string, reverseCorrespondence map[string]string) {
	content := ""
	for i, v := range correspondence {
		content += i + " " + v + "\n"
	}
	if err := os.WriteFile("correspondence"+".txt", []byte(content), 0644); err != nil {
		panic(err)
	}
	content = ""
	for i, v := range reverseCorrespondence {
		content += i + " " + v + "\n"
	}
	if err := os.WriteFile("reverse_correspondence"+".txt", []byte(content), 0644); err != nil {
		panic(err)
	}
}

func Anonymize(baseAddr string, dstAddr string, correspondence map[string]string) {
	for i, v := range correspondence {
		os.MkdirAll(dstAddr, 0766)
		CopyFile(baseAddr+"/"+i, dstAddr+"/"+v)
	}
}
