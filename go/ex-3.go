package main

import (
	"fmt"
	"os"
	"strings"
)

/**
Exercise - 3

3- Replace the word into file and count.

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
	for _, word := range words {
		if strings.TrimRight(word, "\r\n") == findWord {
			os.WriteFile("words-out", []byte(output), 0644)
			println("Replacemt done")
			os.Exit(0)
		}
	}
	println("Word not found")

}
