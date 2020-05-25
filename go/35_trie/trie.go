package main

import "fmt"

type TrieNode struct {
	data         string
	children     []*TrieNode
	isEndingChar bool
}

var (
	root = &TrieNode{
		data:         "/",
		children:     make([]*TrieNode, 26),
		isEndingChar: false,
	}
)

func Insert(text string) {

	p := root

	for i := 0; i < len(text); i++ {
		index := text[i] - 'a'
		if p.children[index] == nil {
			newNode := &TrieNode{
				data:         string(text[i]),
				children:     make([]*TrieNode, 26),
				isEndingChar: false,
			}
			p.children[index] = newNode
		}
		p = p.children[index]
	}

	p.isEndingChar = true
}

func Find(pattern string) bool {

	p := root

	for i := 0; i < len(pattern); i++ {
		index := pattern[i] - 'a'
		if p.children[index] == nil { //不存在pattern
			return false
		}
		p = p.children[index]
	}

	if p.isEndingChar == false { //不完全匹配，只是前缀
		return false
	} else {
		return true
	}
}

func main() {
	Insert("c")
	Insert("bc")
	Insert("bcd")
	Insert("abcd")
	fmt.Println(Find("b"))
	fmt.Println(Find("c"))
	fmt.Println(Find("bd"))
	fmt.Println(Find("abd"))
	fmt.Println(Find("abc"))
	fmt.Println(Find("abcd"))
}
