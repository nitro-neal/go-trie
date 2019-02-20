package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
)

type TrieNode struct {
	character string
	children map[string]TrieNode
	completeWord bool
}

func (node TrieNode) addPhrase(phrase string) {

	if phrase == "" {
		return
	}

	firstChar := string(phrase[0])

	if _, ok := node.children[firstChar]; ok == false {
		if phrase[1:] == "" {
			node.children[firstChar] = TrieNode{firstChar, make(map[string]TrieNode), true}
		} else {
			node.children[firstChar] = TrieNode{firstChar, make(map[string]TrieNode), false}
		}
	}

	node.children[firstChar].addPhrase(phrase[1:])
}

func matchPrefixRecursive(currentNode TrieNode, phrase string) TrieNode {

	fmt.Println(phrase)
	
	if phrase == "" {
		return currentNode
	}

	firstChar := string(phrase[0])

	if _, ok := currentNode.children[firstChar]; ok == false {
		return TrieNode{}
	} else {
		return matchPrefixRecursive(currentNode.children[firstChar], phrase[1:])
	}
}

func autocomplete(currentNode TrieNode, builtString string, initialSearchString string) string {

	if currentNode.completeWord == true {
		fmt.Println("Completed Word: " + initialSearchString + builtString)
	}

	for _, v:= range currentNode.children {
		autocomplete(v, builtString + v.character, initialSearchString)
	}

	return ""
}

func addWordsToTrie(rootNode TrieNode, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		rootNode.addPhrase(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
    fmt.Println("\nExample Trie Implementation in GoLang.\n")
    fmt.Println("This will add all words in a text dictionary to a trie implemtation and retrieve them in an auto complete fashion.")
    fmt.Println("Usage: go 'run exampletrie.go' to use default dictionary path or go 'run exampletrie.go path/to/dictfile' to use custom dictionary\n\n")
    
    var rootNode = TrieNode{"", make(map[string]TrieNode), true}

    // Adds all words in dictionary to trie
	addWordsToTrie(rootNode, "./words_alpha.txt")

	fmt.Println("All words in dictionary added to trie")
}
