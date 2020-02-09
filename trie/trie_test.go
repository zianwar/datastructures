package trie

import (
	"testing"
)

func TestAddWord(t *testing.T) {
	trie := NewTrie()
	trie.AddWord("test")
	expected := true
	got := trie.SearchWord("test")
	if got != expected {
		t.Errorf("expected %v got %v", expected, got)
	}
}
