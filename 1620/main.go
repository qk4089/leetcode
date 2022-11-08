package _1620

import "math"

//给你一个数组towers和一个整数radius。
//数组towers中包含一些网络信号塔，其中towers[i]=[xi,yi,qi]表示第i个网络信号塔的坐标是(xi,yi)且信号强度参数为qi。所有坐标都是在X-Y坐标系内的整数坐标。
//两个坐标之间的距离用欧几里得距离计算。
//
//整数radius表示一个塔能到达的最远距离。如果一个坐标跟塔的距离在radius以内，那么该塔的信号可以到达该坐标。在这个范围以外信号会很微弱，所以radius以外的距离该塔是不能到达的。
//如果第i个塔能到达(x,y)，那么该塔在此处的信号为⌊qi/(1+d)⌋，其中d是塔跟此坐标的距离。一个坐标的信号强度是所有能到达该坐标的塔的信号强度之和。
//
//请你返回数组[cx,cy]，表示信号强度最大的整数坐标点(cx,cy)。如果有多个坐标网络信号一样大，请你返回字典序最小的非负坐标。
//
//注意：
//坐标(x1,y1)字典序比另一个坐标(x2,y2)小，需满足以下条件之一：
//要么x1<x2，
//要么x1==x2且y1<y2。
//⌊val⌋表示小于等于val的最大整数（向下取整函数）

//提示：
//	1 <= towers.length <= 50
//	towers[i].length == 3
//	0 <= xi, yi, qi <= 50
//	1 <= radius <= 50

func main(towers [][]int, radius int) []int {
	if len(towers) == 1 {
		if towers[0][2] != 0 {
			return towers[0][:2]
		} else {
			return []int{0, 0}
		}
	}
	var signals [101][101]int
	maxVal, result := -1, make([]int, 1)
	for _, tower := range towers {
		towerX, towerY := tower[0], tower[1]
		for x := max(towerX-radius, 0); x <= towerX+radius; x++ {
			for y := max(towerY-radius, 0); y <= towerY+radius; y++ {
				d := math.Sqrt(float64((towerX-x)*(towerX-x) + (towerY-y)*(towerY-y)))
				if d <= float64(radius) {
					signals[x][y] += int(float64(tower[2]) / (1 + d))
				}
				if signals[x][y] > maxVal {
					maxVal, result = signals[x][y], []int{x, y}
				} else if signals[x][y] == maxVal && (x < result[0] || (x == result[0] && y < result[1])) {
					result = []int{x, y}
				}
			}
		}
	}
	return result
}

func max(x int, y int) int {
	return int(math.Max(float64(x), float64(y)))
}
