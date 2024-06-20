package _677

//https://leetcode.cn/problems/map-sum-pairs/
//设计一个 map ，满足以下几点: 字符串表示键，整数表示值, 返回具有前缀等于给定字符串的键的值的总和
//实现一个 MapSum 类：
//	MapSum() 初始化 MapSum 对象
//	void insert(String key, int val) 插入key-val键值对，字符串表示键key，整数表示值val。如果键key已经存在，那么原来的键值对 key-value 将被替代成新的键值对。
//	int sum(string prefix) 返回所有以该前缀 prefix 开头的键 key 的值的总和。

//示例 1：
//输入：
//["MapSum", "insert", "sum", "insert", "sum"]
//[[], ["apple", 3], ["ap"], ["app", 2], ["ap"]]
//输出：
//[null, null, 3, null, 5]
//解释：
//MapSum mapSum = new MapSum();
//mapSum.insert("apple", 3);
//mapSum.sum("ap");           // 返回 3 (apple = 3)
//mapSum.insert("app", 2);
//mapSum.sum("ap");           // 返回 5 (apple + app = 3 + 2 = 5)

//提示：
//	1 <= key.length, prefix.length <= 50
//	key 和 prefix 仅由小写英文字母组成
//	1 <= val <= 1000
//	最多调用 50 次 insert 和 sum

type MapSum struct {
	val   int
	child [26]*MapSum
}

func Constructor() MapSum {
	return MapSum{0, [26]*MapSum{}}
}

func (this *MapSum) Insert(key string, val int) {
	point := this
	for _, ch := range key {
		if point.child[ch-'a'] == nil {
			point.child[ch-'a'] = &MapSum{val: 0}
		}
		point = point.child[ch-'a']
	}
	point.val = val
}

func (this *MapSum) Sum(prefix string) int {
	point := this
	for _, ch := range prefix {
		if point.child[ch-'a'] == nil {
			return 0
		}
		point = point.child[ch-'a']
	}
	var postOrder func(dict *MapSum) int
	postOrder = func(dict *MapSum) int {
		if dict == nil {
			return 0
		}
		sum := dict.val
		for i := range dict.child {
			sum += postOrder(dict.child[i])
		}
		return sum
	}
	return postOrder(point)
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
