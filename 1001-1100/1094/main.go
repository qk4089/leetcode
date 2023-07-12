package _1094

//车上最初有capacity个空座位。车只能向一个方向行驶（也就是说，不允许掉头或改变方向）
//给定整数capacity和一个数组trips,trip[i]=[numPassengersi, fromi, toi]表示第i次旅行有numPassengersi乘客，接他们和放他们的位置分别是fromi和toi。
//这些位置是从汽车的初始位置向东的公里数。 当且仅当你可以在所有给定的行程中接送所有乘客时，返回true，否则请返回false。

//提示：
//	1<=trips.length<=1000
//	trips[i].length==3
//	1<=numPassengersi<=100
//	0<=fromi<toi<=1000
//	1<=capacity<=105

//示例
//输入：trips=[[2,1,5],[3,3,7]],capacity=4
//输出：false

//输入：trips=[[2,1,5],[3,3,7]],capacity=5
//输出：true

func carPooling(trips [][]int, capacity int) bool {
	distance := make([]int, 1000)
	for _, trip := range trips {
		for i := trip[1]; i <= trip[2]; i++ {
			distance[i] += trip[0]
			if distance[i] > capacity {
				return false
			}
		}
	}
	return true
}
