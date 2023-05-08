package _1109

//这里有n个航班，它们分别从1到n进行编号。
//有一份航班预订表bookings，表中第i条预订记录bookings[i]=[firsti,lasti,seatsi]意味着在从firsti到lasti（包含firsti和lasti）的每个航班上预订了seatsi个座位。
//请你返回一个长度为n的数组answer，里面的元素是每个航班预定的座位总数。

//提示：
//	1<=n<=2*104
//	1<=bookings.length<=2*104
//	bookings[i].length==3
//	1<=firsti<=lasti<=n
//	1<=seatsi<=104

//示例
//输入：bookings=[[1,2,10],[2,3,20],[2,5,25]],n=5
//输出：[10,55,45,25,25]
//解释：
//	航班编号		1	2	3	4	5
//	预订记录1：  10  10
//	预订记录2：      20  20
//	预订记录3：      25  25  25  25
//	总座位数：   10  55  45  25  25
//	因此，answer=[10,55,45,25,25]

//输入：bookings=[[1,2,10],[2,2,15]],n=2
//输出：[10,25]
//解释：
//	航班编号12
//	预订记录1：1010
//	预订记录2：15
//	总座位数：1025
//	因此，answer=[10,25]

func corpFlightBookings(bookings [][]int, n int) []int {
	diff := make([]int, n)
	for _, booking := range bookings {
		left, right, val := booking[0], booking[1], booking[2]
		diff[left-1] += val
		if right+1 < n {
			diff[right+1] -= val
		}
	}
	return restore(diff)
}

func restore(diff []int) []int {
	for i := 1; i < len(diff); i++ {
		diff[i] = diff[i-1] + diff[i]
	}
	return diff
}
