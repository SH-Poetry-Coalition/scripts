package main

import (
	"fmt"
	"os"
	"scripts/util"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Expected 'archive' or 'anonymize' or 'reverse' subcommands!")
	} else {
		switch os.Args[1] {
		case "archive":
			poems := util.ReadPoems("./poems_original")
			util.Archive("./poems_original", "./poems", poems)
		case "anonymize":
			poems := util.ReadPoems("./poems_original")
			correspondence, reverseCorrespondence := util.GenCorrespondence("./poems_original", poems)
			util.WriteCorrespondence(correspondence, reverseCorrespondence)
			util.Anonymize("./poems_original", "./poems_anonymous", correspondence)
		case "reverse":
			reverseCorrespondence := util.ReadCorrespondence("reverse_correspondence.txt")
			util.Reverse("./poems_anonymous", "./poems_reversed", reverseCorrespondence)
		default:
			fmt.Println("Expected 'archive' or 'anonymize' or 'reverse' subcommands!")
		}
	}
}
