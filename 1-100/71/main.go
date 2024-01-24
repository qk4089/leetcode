package _71

import (
	"bytes"
	"strings"
)

//https://leetcode.cn/problems/simplify-path/
//给你一个字符串path，表示指向某一文件或目录的Unix风格绝对路径（以'/'开头），请你将其转化为更加简洁的规范路径。
//在Unix风格的文件系统中，一个点（.）表示当前目录本身；此外，两个点（..）表示将目录切换到上一级（指向父目录）；两者都可以是复杂相对路径的组成部分。
//对于此问题，任意多个连续的斜杠（即，'//'）都被视为单个斜杠'/'。任何其他格式的点（例如，'...'）均被视为文件/目录名称。

//请注意，返回的规范路径必须遵循下述格式：
//	始终以斜杠'/'开头。
//	两个目录名之间必须只有一个斜杠'/'。
//	最后一个目录名（如果存在）不能以'/'结尾。
//	此外，路径仅包含从根目录到目标文件或目录的路径上的目录（即，不含'.'或'..'）。

//返回简化后得到的规范路径。

//示例1：
//输入：path="/home/"
//输出："/home"
//解释：注意，最后一个目录名后面没有斜杠。

//输入：path="/../"
//输出："/"
//解释：从根目录向上一级是不可行的，因为根目录是你可以到达的最高级。

//输入：path="/home//foo/"
//输出："/home/foo"
//解释：在规范路径中，多个连续斜杠需要用一个斜杠替换。

//输入：path="/a/./b/../../c/"
//输出："/c"

//提示：
//	1<=path.length<=3000
//	path由英文字母，数字，'.'，'/'或'_'组成。
//	path是一个有效的Unix风格绝对路径。

func simplifyPath(path string) string {
	ans, segment := make([][]byte, 0), make([]byte, 0)
	for i := 0; i < len(path); i++ {
		switch path[i] {
		case '/':
			if len(segment) > 0 {
				ans, segment = append(ans, append([]byte{}, segment...)), make([]byte, 0)
			}
		case '.':
			if len(segment) > 0 {
				segment = append(segment, '.')
				continue
			}
			cnt := 0
			for ; i < len(path) && path[i] == '.'; cnt++ {
				i++
			}
			if cnt >= 3 || (i < len(path) && path[i] != '/') {
				segment = append(segment, strings.Repeat(".", cnt)...)
				i--
				continue
			}
			if cnt == 2 && len(ans) > 0 {
				ans = ans[:len(ans)-1]
			}
		default:
			segment = append(segment, path[i])
		}
	}
	if len(segment) > 0 {
		ans = append(ans, append([]byte{}, segment...))
	}
	return "/" + string(bytes.Join(ans, []byte{'/'}))
}
