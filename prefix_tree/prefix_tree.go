package prefix_tree

type Trie struct {
	children [26]*Trie
	isLeaf   bool
}

func Constructor() Trie {
	trie := Trie{}

	trie.isLeaf = false
	return trie
}

func (this *Trie) Insert(word string) {
	cur := this

	for _, ch := range word {
		idx := ch - 'a'

		if cur.children[idx] == nil {
			cur.children[idx] = &Trie{}
		}

		cur = cur.children[idx]
	}

	cur.isLeaf = true
}

func (this *Trie) Search(word string) bool {
	cur := this

	for _, ch := range word {
		idx := ch - 'a'

		if cur.children[idx] == nil {
			return false
		}

		cur = cur.children[idx]
	}

	return cur.isLeaf
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this

	for _, ch := range prefix {
		idx := ch - 'a'

		if cur.children[idx] == nil {
			return false
		}

		cur = cur.children[idx]
	}

	return true
}
