package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var punctuationRegex = regexp.MustCompile(`[^\w\s]`)

func countWords(text string) int {
	cleanText := punctuationRegex.ReplaceAllString(text, "")
	words := strings.Fields(cleanText)

	return len(words)
}

func main() {
	fmt.Println("Enter your phrase:")

	stdinReader := bufio.NewReader(os.Stdin)
	userInput, err := stdinReader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	wordCount := countWords(userInput)

	fmt.Println("Word count: ", wordCount)
}
