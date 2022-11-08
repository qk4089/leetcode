package _481

//神奇字符串s仅由'1'和'2'组成，并需要遵守下面的规则：
//
//神奇字符串s的神奇之处在于，串联字符串中'1'和'2'的连续出现次数可以生成该字符串。
//s的前几个元素是s="1221121221221121122……"。
//如果将s中连续的若干1和2进行分组，可以得到"1221121221221121122......"。
//每组中1或者2的出现次数分别是"122112122122......"。上面的出现次数正是s自身。
//给你一个整数n，返回在神奇字符串s的前n个数字中1的数目。
//提示：
//	1 <= n <= 105

//示例
//输入：n = 6
//输出：3
//解释：神奇字符串 s 的前 6 个元素是 “122112”，它包含三个 1，因此返回 3

func main(n int) int {
	count := 1
	if n < 4 {
		return count
	}
	point := 2
	var last byte = 2
	src := make([]byte, 0, n+2)
	src = append(src, 1, 2, 2)
	nextMap := map[byte]byte{1: 2, 2: 1}
	for len(src) < n {
		num := src[point]
		next := nextMap[last]
		for i := 0; i < int(num); i++ {
			if next == 1 {
				count++
			}
			src = append(src, next)
		}
		point++
		last = next
	}
	if len(src) > n && src[len(src)-1:][0] == 1 {
		count--
	}
	return count
}
