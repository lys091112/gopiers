package base

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := TireConstructor()

	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))
	fmt.Println(trie.Search("app"))
	fmt.Println(trie.StartsWith("app"))
	trie.Insert("app")
	fmt.Println(trie.Search("app"))
	fmt.Println(trie.StartsWith("app"))
}
