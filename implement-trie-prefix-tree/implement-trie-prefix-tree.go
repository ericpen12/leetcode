package main

type Trie struct {
	Child []Trie
	isEnd bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{Child: make([]Trie, 26)}
}

func (this *Trie) Put(ch int, node Trie) {
	this.Child[ch] = node
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	for _, v := range word {
		if this.Child[v-'a'].Child == nil {
			this.Put(int(v-'a'), Trie{Child: make([]Trie, 26)})
		}
		this = &this.Child[v-'a']
	}
	this.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	return false
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

func main() {
	word := "hello"

	obj := Constructor()
	obj.Insert(word)
}
