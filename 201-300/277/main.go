package _277

//给你 n 个人的社交关系（你知道任意两个人之间是否认识），然后请你找出这些人中的「名人」。
//所谓「名人」有两个条件：
//	1、所有其他人都认识「名人」。
//	2、「名人」不认识任何其他人。

//示例 1：
//输入：graph = [[1,1,0],[0,1,0],[1,1,1]]
//输出：1

//在本题中，你可以使用辅助函数 bool knows(a, b) 获取到 A 是否认识 B。请你来实现一个函数 int findCelebrity(n)。
//派对最多只会有一个 “名人” 参加。若 “名人” 存在，请返回他/她的编号；若 “名人” 不存在，请返回 -1。

func findCelebrity(graph [][]int) int {
	inputMap := make(map[int]int)
	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph); j++ {
			if i == j {
				continue
			}
			inputMap[j] += graph[i][j]
		}
	}
	for i, v := range inputMap {
		if v == len(graph)-1 {
			for j := 0; j < len(graph[i]); j++ {
				if i == j {
					continue
				}
				if graph[i][j] != 0 {
					return -1
				}
			}
			return 1
		}
	}
	return -1
}
