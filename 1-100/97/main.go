package _97

//https://leetcode.cn/problems/interleaving-string/
//给定三个字符串 s1、s2、s3，请你帮忙验证 s3 是否是由 s1 和 s2 交错 组成的。

//两个字符串 s 和 t 交错 的定义与过程如下，其中每个字符串都会被分割成若干 非空 子字符串：
//s = s1 + s2 + ... + sn, t = t1 + t2 + ... + tm, |n - m| <= 1
//交错 是 s1 + t1 + s2 + t2 + s3 + t3 + ... 或者 t1 + s1 + t2 + s2 + t3 + s3 + ...
//注意：a + b 意味着字符串 a 和 b 连接。

//示例 1：
//输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
//输出：true

//输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
//输出：false

//输入：s1 = "", s2 = "", s3 = ""
//输出：true

//提示：
//	0 <= s1.length, s2.length <= 100
//	0 <= s3.length <= 200
//	s1、s2、和 s3 都由小写英文字母组成

//进阶：您能否仅使用 O(s2.length) 额外的内存空间来解决它?

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	dp := make([][]bool, len(s1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s2)+1)
	}
	dp[0][0] = true
	for i := 0; i <= len(s1); i++ {
		for j := 0; j <= len(s2); j++ {
			if i > 0 {
				dp[i][j] = dp[i][j] || s1[i-1] == s3[i-1+j] && dp[i-1][j]
			}
			if j > 0 {
				dp[i][j] = dp[i][j] || s2[j-1] == s3[i+j-1] && dp[i][j-1]
			}
		}
	}
	return dp[len(s1)][len(s2)]
	//return backtrace(0, 0, 0, s1, s2, s3)
}

func backtrace(p1, p2, p3 int, s1 string, s2 string, s3 string) bool {
	if p3 == len(s3) {
		return true
	}
	if p1 == len(s1) {
		return s2[p2:] == s3[p3:]
	}
	if p2 == len(s2) {
		return s1[p1:] == s3[p3:]
	}
	if s1[p1] == s3[p3] && backtrace(p1+1, p2, p3+1, s1, s2, s3) {
		return true
	} else if s2[p2] == s3[p3] && backtrace(p1, p2+1, p3+1, s1, s2, s3) {
		return true
	}
	return false
}
