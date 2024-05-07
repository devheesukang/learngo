package main

import (
	"fmt"

	mydict "github.com/devheesukang/learngo/go-dictionary/dict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	word := "hello"
	definition := "Greeting"

	// Add new word
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
	// Search for added word
	hello, _ := dictionary.Search(word)
	fmt.Println("found", word, "definition:", hello)

	// Add word that already exists
	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Println(err2)
	}

	// Update word
	err3 := dictionary.Update(word, "Second")
	if err3 != nil {
		fmt.Println(err3)
	}
	baseword, _ := dictionary.Search(word)
	fmt.Println(baseword)

	// Update non-existing word
	err4 := dictionary.Update("random", "Second")
	if err4 != nil {
		fmt.Println(err4)
	} else {
		baseword2, _ := dictionary.Search("random")
		fmt.Println(baseword2)
	}

	// Delete word
	dictionary.Delete(word)
	word, err5 := dictionary.Search(word)
	if err5 != nil {
		fmt.Println(err5)
	} else {
		fmt.Println(word)
	}
}
