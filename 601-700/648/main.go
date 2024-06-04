package _648

import "strings"

//https://leetcode.cn/problems/replace-words/
//在英语中，我们有一个叫做 词根(root) 的概念，可以词根后面添加其他一些词组成另一个较长的单词——我们称这个词为继承词 (successor)。
//例如，词根 help，跟随着继承词 "ful"，可以形成新的单词 "helpful"。

//现在，给定一个由许多词根组成的词典dictionary和一个用空格分隔单词形成的句子sentence。你需要将句子中的所有继承词用词根替换掉。
//如果继承词有许多可以形成它的词根，则用最短的词根替换它。你需要输出替换之后的句子。

//示例 1：
//输入：dictionary = ["cat","bat","rat"], sentence = "the cattle was rattled by the battery"
//输出："the cat was rat by the bat"

//输入：dictionary = ["a","b","c"], sentence = "aadsfasf absbs bbab cadsfafs"
//输出："a a b c"

//提示：
//	1 <= dictionary.length <= 1000
//	1 <= dictionary[i].length <= 100
//	dictionary[i] 仅由小写字母组成。
//	1 <= sentence.length <= 106
//	sentence 仅由小写字母和空格组成。
//	sentence 中单词的总量在范围 [1, 1000] 内。
//	sentence 中每个单词的长度在范围 [1, 1000] 内。
//	sentence 中单词之间由一个空格隔开。
//	sentence 没有前导或尾随空格。

func replaceWords(dictionary []string, sentence string) string {
	tire, words := Constructor(), strings.Split(sentence, " ")
	for _, dict := range dictionary {
		tire.Add(dict)
	}
	ans := make([]string, len(words))
	for i, word := range words {
		ans[i] = tire.Get(word)
	}
	return strings.Join(ans, " ")
}

type Trie struct {
	isWord   bool
	children [26]*Trie
}

func Constructor() *Trie {
	return &Trie{false, [26]*Trie{}}
}

func (this *Trie) Add(word string) {
	point := this
	for _, char := range word {
		if point.children[char-'a'] == nil {
			point.children[char-'a'] = Constructor()
		}
		point = point.children[char-'a']
	}
	point.isWord = true
}

func (this *Trie) Get(word string) string {
	point, ans := this, make([]byte, 0)
	for i := 0; i < len(word); i++ {
		if point.children[word[i]-'a'] == nil || point.isWord {
			break
		}
		ans, point = append(ans, word[i]), point.children[word[i]-'a']
	}
	if len(ans) > 0 && point.isWord {
		return string(ans)
	} else {
		return word
	}
}
