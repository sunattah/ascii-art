package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := os.Args[2]
	if input == " " {
		fmt.Println("Error")
	}

}

func LoadBanner(filename string) (map[rune][]string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(file), "\n")

	bannerMap := make(map[rune][]string)

	currentChar := rune(32)

	for i := 1; i < len(lines); i += 9 {

		if i+8 > len(lines) {

			break
		}

		bannerMap[currentChar] = lines[i : i+8]

		currentChar++
	}

	return bannerMap, nil
}

func ParseInput(input string) []string {

	words := strings.Split(strings.ReplaceAll(input, "\\n", "\n"), "\n")

	return words
}

func PrintASCII(words []string, banner map[rune][]string) {
	for _, word := range words {
		if word == "" {
			fmt.Println()
			continue
		}

		for row := 0; row < 8; row++ {
			for _, char := range word {
				charArt := banner[char]

				fmt.Print(charArt[row])
			}
			fmt.Println()
		}
	}
}

func IsValid(input string) bool {
	lines := strings.Split(input, " ")
	for _, ch := range lines {

		if ch == LoadBanner(input) {

		}

	}
	return false
}
