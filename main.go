package main

import (
	"bufio"
	"fmt"
	"myDictionary/dictionary"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	d := dictionary.New()
	for {
		fmt.Println("Chose a action: [ADD/DEFINE/REMOVE/LIST/EXIT]:")
		userCommand, _ := reader.ReadString('\n')
		userCommand = strings.TrimSpace(userCommand)

		switch userCommand {
		case "ADD":
			actionAdd(d, reader)
		case "DEFINE":
			actionDefine(d, reader)
		case "REMOVE":
			actionRemove(d, reader)
		case "LIST":
			actionList(d)
		case "EXIT":
			fmt.Println("End of program")
			return
		default:
			fmt.Println("Unknown action")
		}
	}
}

func actionAdd(d *dictionary.Dictionary, reader *bufio.Reader) {
	fmt.Println("Enter a word:")
	newWord, _ := reader.ReadString('\n')
	newWord = strings.TrimSpace(newWord)
	fmt.Println("Enter a definition")
	definition, _ := reader.ReadString('\n')
	definition = strings.TrimSpace(definition)
	d.Add(newWord, definition)
	fmt.Println("Word added succesfully")
}

func actionDefine(d *dictionary.Dictionary, reader *bufio.Reader) {
	fmt.Println("Enter a word:")
	word, _ := reader.ReadString('\n')
	word = strings.TrimSpace(word)

	entry, err := d.Get(word)
	if err != nil {
		fmt.Println("Word not found.")
	} else {
		fmt.Println("Definition:", entry.String())
	}
}

func actionRemove(d *dictionary.Dictionary, reader *bufio.Reader) {
	fmt.Println("Enter a word:")
	wordToDelete, _ := reader.ReadString('\n')
	wordToDelete = strings.TrimSpace(wordToDelete)
	d.Remove(wordToDelete)
}

func actionList(d *dictionary.Dictionary) {
	words, entries := d.List()
	for _, word := range words {
		fmt.Println(word, ":", entries[word].String())
	}
}
