package base

type Trie struct {
	value  string
	isEnd  bool
	childs map[string]*Trie
	nums   int
}

// TODO 当childs的包含对象为Trie而非*Trie时，貌似会无法修改isEnd参数
/** Initialize your data structure here. */
func TireConstructor() Trie {
	return Trie{
		value:  "",
		isEnd:  false,
		childs: make(map[string]*Trie),
		nums:   0}

}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	if word == "" {
		return
	}

	ws := []rune(word)

	tmp := this
	for i := 0; i < len(ws); i++ {
		key := string(ws[i])
		(*tmp).nums += 1
		if c, ok := tmp.childs[key]; ok {
			tmp = c
		} else {
			newTrie := TireConstructor()
			newTrie.value = key
			tmp.childs[key] = &newTrie
			tmp = &newTrie
		}
	}
	tmp.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	t := this.findEndNode(word)
	return t != nil && t.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	return this.findEndNode(prefix) != nil
}

func (this *Trie) findEndNode(word string) *Trie {
	if word == "" {
		return nil
	}

	ws := []rune(word)

	tmp := this
	for _, v := range ws {
		key := string(v)
		if c, ok := tmp.childs[key]; ok {
			tmp = c
		} else {
			return nil
		}
	}
	return tmp

}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
