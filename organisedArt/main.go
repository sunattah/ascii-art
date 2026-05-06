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
	input := os.Args[1]

	if !IsValid(input) {
		fmt.Println("Error: Invalid characters in input")
		return
	}

	banner, err := LoadBanner("thinkertoy.txt")
	if err != nil {
		fmt.Println("Error loading banner:", err)
		return
	}

	words := ParseInput(input)

	PrintASCII(words, banner)
}

// func LoadBanner(filename string) (map[rune][]string, error) {
// 	file, err := os.ReadFile(filename)
// 	if err != nil {
// 		return nil, err
// 	}

// 	lines := strings.Split(string(file), "\n")

// 	bannerMap := make(map[rune][]string)

// 	currentChar := rune(32)

// 	for i := 1; i < len(lines); i += 9 {

// 		if i+8 > len(lines) {

// 			break
// 		}

// 		bannerMap[currentChar] = lines[i : i+8]

// 		currentChar++
// 	}

// 	return bannerMap, nil
// }

func ParseInput(input string) []string {

	words := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\r\n")

	return words
}

// func PrintASCII(words []string, banner map[rune][]string) {
// 	for _, word := range words {
// 		if word == "" {
// 			fmt.Println()
// 			continue
// 		}

// 		for _, word := range words {
// 			for row := 0; row < 8; row++ {
// 				for _, char := range word {
// 					charArt := banner[char]

// 					fmt.Print(charArt[row])
// 				}
// 				fmt.Println()
// 			}
// 		}
// 	}
// }

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

    content := strings.ReplaceAll(string(file), "\r\n", "\n")

    lines := strings.Split(content, "\n")

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


func PrintASCII(words []string, banner map[rune][]string) {
    for _, word := range words {
        if word == "" {
            fmt.Println()
            continue
        }

        for row := 0; row < 8; row++ {
            for _, char := range word {
                charArt, exists := banner[char]
                if exists {
                    fmt.Print(charArt[row])
                }
            }
            fmt.Println()
        }
    }
}

