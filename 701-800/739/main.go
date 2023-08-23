package _739

//给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指对于第 i 天，下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用 0 来代替。
//
//示例 1:
//输入: temperatures = [73,74,75,71,69,72,76,73]
//输出: [1,1,4,2,1,1,0,0]

//输入: temperatures = [30,40,50,60]
//输出: [1,1,1,0]

//输入: temperatures = [30,60,90]
//输出: [1,1,0]

//提示：
//	1 <= temperatures.length <= 10^5
//	30 <= temperatures[i] <= 100

func dailyTemperatures(temperatures []int) []int {
	stack, result := make([]int, 0), make([]int, len(temperatures))
	for i := len(temperatures) - 1; i >= 0; i-- {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] <= temperatures[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			result[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
	}
	return result
}
