package trie

type Node struct {
	Children map[rune]*Node
	IsEnd    bool
	Freq     int
	Word     string
}
