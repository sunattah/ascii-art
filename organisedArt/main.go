package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	inputText := os.Args
	file, err := os.ReadFile(inputText[2])
	if err != nil {
		fmt.Println("Error while reading file")
		return
	}

	lines := strings.Split(string(file), "\n")
	word := strings.Split(strings.ReplaceAll(inputText[1], "\\n", "\n"), "\n")
	//replace := strings.ReplaceAll(lines[file[]], "\r\n", "\n")
	for _, text := range word {
		if text == "" {
			fmt.Println()
		}

		for start := 1; start <= 8; start++ {
			for _, char := range text {
				star := int(char-32)*9 + 1

				fmt.Print(lines[star+start-1])
			}
			fmt.Println()
		}
	}

}
