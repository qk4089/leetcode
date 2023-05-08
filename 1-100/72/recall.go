package _72

func recall(word1, word2 string) int {
	return execute(word1, word2, 0, 0, 0)
}

func execute(word1, word2 string, i, j, dist int) int {
	if i == len(word1) || j == len(word2) {
		return dist + len(word1) - i + len(word2) - j
	}
	if word1[i] == word2[j] {
		return execute(word1, word2, i+1, j+1, dist)
	} else {
		return min(execute(word1, word2, i+1, j+1, dist+1), min(execute(word1, word2, i, j+1, dist+1), execute(word1, word2, i+1, j, dist+1)))
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
