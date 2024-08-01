package trie

type Trie struct {
	children [26]*Trie
	end      bool
}

func NewTrie() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	curr := t
	for _, ch := range word {
		ch -= 'a'
		if curr.children[ch] == nil {
			curr.children[ch] = &Trie{}
		}
		curr = curr.children[ch]
	}
	curr.end = true
}
func (t *Trie) search(word string) *Trie {
	curr := t
	for _, ch := range word {
		ch -= 'a'
		if curr.children[ch] == nil {
			return nil
		}
		curr = curr.children[ch]
	}
	return curr
}

func (t *Trie) Search(word string) bool {
	curr := t.search(word)
	return curr != nil && curr.end
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.search(prefix) != nil
}
