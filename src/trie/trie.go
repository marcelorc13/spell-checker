package trie

type Trie struct {
	root *Node
}

func New() *Trie {
	return &Trie{root: &Node{Children: make(map[rune]*Node)}}
}

func (t *Trie) Root() *Node {
	return t.root
}

func (t *Trie) Insert(word string, freq int) {
	cur := t.root
	for _, ch := range word {
		if cur.Children[ch] == nil {
			cur.Children[ch] = &Node{Children: make(map[rune]*Node)}
		}
		cur = cur.Children[ch]
	}
	cur.IsEnd = true
	cur.Freq = freq
	cur.Word = word
}

func (t *Trie) Search(word string) (freq int, found bool) {
	cur := t.root
	for _, ch := range word {
		if cur.Children[ch] == nil {
			return 0, false
		}
		cur = cur.Children[ch]
	}
	if cur.IsEnd {
		return cur.Freq, true
	}
	return 0, false
}
