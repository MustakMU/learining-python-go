package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("chat.json")
	reader := bufio.NewReader(os.Stdin)

	check(err)
	content := string(dat)

	var chatData map[string]interface{}

	err = json.Unmarshal([]byte(content), &chatData)
	check(err)
	fmt.Println(chatData)

	println("Hi.. How i can help you?")
	for {
		cmd, _ := reader.ReadString('\n')

		if cmd == "exit" || cmd == "quit" {
			fmt.Println("Take care ... Bye!")
		}
		cmd = strings.TrimSuffix(cmd, "\n")
		if len(strings.Split(cmd, " ")) < 2 {
			fmt.Println("Please enter atleast two words")
			continue
		} else {
			cmdSplit := strings.Split(cmd, " ")
			indexes := map[string]int{}
			highestIndex := 0
			for key, _ := range chatData {
				matches := 0
				keySplit := strings.Split(key, " ")
				for _, word := range cmdSplit {
					for _, resp := range keySplit {
						if word == resp {
							matches++
						}
					}
				}
				if highestIndex < matches {
					highestIndex = matches
				}
				indexes[key] = matches
			}
			if highestIndex == 0 {
				fmt.Println("Command not found! please try again")
			} else {
				for key, val := range indexes {
					if val == highestIndex {
						innermap, _ := chatData[key].(map[string]interface{})
						fmt.Println(innermap["reply"])
						break
					}
				}
			}
		}

	}
}
