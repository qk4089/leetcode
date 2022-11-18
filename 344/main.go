package _344

//编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组s的形式给出。
//不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用O(1)的额外空间解决这一问题。

//提示：
//1<=s.length<=105
//s[i]都是ASCII码表中的可打印字符

//示例
//输入：s=["h","e","l","l","o"]
//输出：["o","l","l","e","h"]

//输入：s=["H","a","n","n","a","h"]
//输出：["h","a","n","n","a","H"]

func reverseString(s []byte) {
	if len(s) == 1 {
		return
	}
	start, end := 0, len(s)-1
	for end-start > 2 {
		s[start], s[end], start, end = s[end], s[start], start+1, end-1
	}
	s[start], s[end] = s[end], s[start]
}
