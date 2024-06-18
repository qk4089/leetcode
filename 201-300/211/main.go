package _211

//https://leetcode.cn/problems/design-add-and-search-words-data-structure/
//请你设计一个数据结构，支持添加新单词和查找字符串是否与任何先前添加的字符串匹配 。

//实现词典类 WordDictionary ：
//	WordDictionary() 初始化词典对象
//	void addWord(word) 将word添加到数据结构中，之后可以对它进行匹配
//	bool search(word) 如果数据结构中存在字符串与word匹配，则返回true；否则返回false。word中可能包含一些'.'，每个.都可以表示任何一个字母。

//示例：
//输入：
//["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
//[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
//输出：
//[null,null,null,null,false,true,true,true]
//解释：
//	WordDictionary wordDictionary = new WordDictionary();
//	wordDictionary.addWord("bad");
//	wordDictionary.addWord("dad");
//	wordDictionary.addWord("mad");
//	wordDictionary.search("pad"); // 返回 False
//	wordDictionary.search("bad"); // 返回 True
//	wordDictionary.search(".ad"); // 返回 True
//	wordDictionary.search("b.."); // 返回 True

//提示：
//	1 <= word.length <= 25
//	addWord 中的 word 由小写英文字母组成
//	search 中的 word 由 '.' 或小写英文字母组成
//	最多调用 10^4 次 addWord 和 search

type WordDictionary struct {
	isWord bool
	child  [26]*WordDictionary
}

func Constructor() WordDictionary {
	return WordDictionary{false, [26]*WordDictionary{}}
}

func (this *WordDictionary) AddWord(word string) {
	point := this
	for _, ch := range word {
		if point.child[ch-'a'] == nil {
			point.child[ch-'a'] = &WordDictionary{}
		}
		point = point.child[ch-'a']
	}
	point.isWord = true
}

func (this *WordDictionary) Search(word string) bool {
	return backtrace(this, word)
}

func backtrace(dict *WordDictionary, word string) bool {
	if len(word) == 0 {
		return dict.isWord
	}
	if word[0] == '.' {
		for i := 0; i < 26; i++ {
			if dict.child[i] != nil && backtrace(dict.child[i], word[1:]) {
				return true
			}
		}
		return false
	} else {
		if dict.child[word[0]-'a'] == nil {
			return false
		}
		return backtrace(dict.child[word[0]-'a'], word[1:])
	}
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
