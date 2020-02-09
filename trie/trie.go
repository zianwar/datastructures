package trie

type TrieNode struct {
	isWord   bool
	char     rune
	children map[rune]*TrieNode
}

func NewTrieNode(char rune) *TrieNode {
	return &TrieNode{
		char:     char,
		children: make(map[rune]*TrieNode),
	}
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(rune(0)),
	}
}

func (t *Trie) AddWord(word string) {
	current := t.root

	for _, c := range word {
		if next, ok := current.children[c]; ok {
			current = next
		} else {
			tmp := NewTrieNode(c)
			current.children[c] = tmp
			current = tmp
		}
	}

	current.isWord = true
}

func (t *Trie) SearchWord(word string) bool {
	current := t.root

	for _, c := range word {
		if next, ok := current.children[c]; ok {
			current = next
		} else {
			return false
		}
	}

	return current.isWord
}
