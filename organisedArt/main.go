package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		return
	}
	input := os.Args

	if input[1] == "" {
		return
	}


	if !IsValid(input[1]) {
		fmt.Println("Error: Invalid characters in input")
		return
	}

	file := "standard.txt"

	if len(os.Args) > 2 {
		file = os.Args[2]
		if !strings.HasSuffix(file, ".txt") {
			file = file+".txt"
		}
	}

	banner, err := LoadBanner(file)
	if err != nil {
		fmt.Println("Error loading banner:", err)
		return
	}
	words := ParseInput(input[1])

	PrintASCII(words, banner)
}

func ParseInput(input string) []string {
	text := strings.Split(strings.ReplaceAll(input, "\\n", "\n"), "\n")

	return text

}

func IsValid(input string) bool {
	for _, ch := range input {
		if ch != 10 && (ch < 32 || ch > 126) {
			return false
		}
	}
	return true
}
func LoadBanner(filename string) (map[rune][]string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	content := strings.Split(strings.ReplaceAll(string(file), "\r\n", "\n"), "\n")

	bannerMap := make(map[rune][]string)
	currentChar := rune(32)

	for i := 1; i < len(content); i += 9 {
		if i+8 > len(content) {
			break
		}
		bannerMap[currentChar] = content[i : i+8]
		currentChar++
	}
	if len(bannerMap) != 95 {
		return nil, fmt.Errorf("expected 95 characters, got %d", len(bannerMap))
	}

	return bannerMap, nil
}


func PrintASCII(words []string, banner map[rune][]string) {
	for i, word := range words {
		if word == "" {
			// Only skip the very last empty string produced by Split
			if i < len(words)-1 {
				fmt.Println()
			}
			continue
		}

		for row := 0; row < 8; row++ {
			for _, char := range word {
				fmt.Print(banner[char][row])
			}
			fmt.Println()
		}
	}
}
