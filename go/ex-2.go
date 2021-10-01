package main

import (
	"fmt"
	"os"
	"strings"
)

/**
Exercise - 2

2- Search and count the word available in file.
**/

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("words.txt")
	check(err)
	content := string(dat)

	println("Enter word to replace: ")
	var findWord string
	fmt.Scanln(&findWord)

	println("Enter replacement word: ")
	var replacementWord string
	fmt.Scanln(&replacementWord)
	output := ""
	words := strings.Split(content, " ")
	for i, word := range words {
		if strings.Contains(word, findWord) {
			words[i] = strings.Replace(word, findWord, replacementWord, 1)
		} else {
			output += word
		}
	}
	for i := 1; i < len(words); i++ {
		println("words[i] " + words[i])
		output += " " + words[i]
	}
	os.WriteFile("words-out", []byte(output), 0644)
}
