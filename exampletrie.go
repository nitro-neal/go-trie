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
	fmt.Println("Will add phrase to trie: ", phrase)
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
    fmt.Println("Example Trie Implementation")
    var rootNode = TrieNode{"", make(map[string]TrieNode), true}

    // Adds all words in dictionary to trie
	addWordsToTrie(rootNode, "./words_alpha.txt")
}
