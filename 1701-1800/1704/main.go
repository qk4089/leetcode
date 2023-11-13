package _1704

//给你一个偶数长度的字符串s。将其拆分成长度相同的两半，前一半为a，后一半为b。
//两个字符串相似的前提是它们都含有相同数目的元音（'a'，'e'，'i'，'o'，'u'，'A'，'E'，'I'，'O'，'U'）。注意，s可能同时含有大写和小写字母。
//如果a和b相似，返回true；否则，返回false。

//提示：
//	2<=s.length<=1000
//	s.length是偶数
//	s由大写和小写字母组成

//示例
//输入：s="book"
//输出：true
//解释：a="bo"且b="ok"。a中有1个元音，b也有1个元音。所以，a和b相似。

//输入：s="textbook"
//输出：false
//解释：a="text"且b="book"。a中有1个元音，b中有2个元音。因此，a和b不相似。注意，元音o在b中出现两次，记为2个

func halvesAreAlike(s string) bool {
	meta := map[byte]byte{'a': 1, 'e': 1, 'i': 1, 'o': 1, 'u': 1, 'A': 1, 'E': 1, 'I': 1, 'O': 1, 'U': 1}
	p, c1, c2 := len(s)/2-1, 0, 0
	for i := p + 1; i < len(s); i++ {
		if _, ok := meta[s[i]]; ok {
			c1++
		}
		if _, ok := meta[s[p]]; ok {
			c2++
		}
		p--
	}
	return c1 == c2
}
