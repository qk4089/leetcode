package str

type Trie struct {
	child map[byte]*Trie
	flag  bool
}

func NewTire() *Trie {
	return &Trie{child: make(map[byte]*Trie), flag: false}
}

func (t *Trie) insert(str string) {
	root := t
	for i := 0; i < len(str); i++ {
		if _, ok := root.child[str[i]]; !ok {
			root.child[str[i]] = &Trie{child: make(map[byte]*Trie), flag: false}
		}
		root = root.child[str[i]]
	}
	root.flag = true
}

func (t *Trie) find(str string) bool {
	root := t
	for i := 0; i < len(str); i++ {
		if val, ok := root.child[str[i]]; ok {
			root = val
		} else {
			return false
		}
	}
	return root.flag
}
