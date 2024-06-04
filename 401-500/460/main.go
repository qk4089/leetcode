package _460

//https://leetcode.cn/problems/lfu-cache/
//请你为 最不经常使用（LFU）缓存算法设计并实现数据结构。
//实现 LFUCache 类：
//	LFUCache(int capacity) - 用数据结构的容量 capacity 初始化对象
//	int get(int key) - 如果键 key 存在于缓存中，则获取键的值，否则返回 -1 。
//	void put(int key, int value) - 如果键 key 已存在，则变更其值；如果键不存在，请插入键值对。当缓存达到其容量 capacity 时，则应该在插入新项之前，移除最不经常使用的项。在此问题中，当存在平局（即两个或更多个键具有相同使用频率）时，应该去除 最久未使用 的键。

//为了确定最不常使用的键，可以为缓存中的每个键维护一个使用计数器。使用计数最小的键是最久未使用的键。当一个键首次插入到缓存中时，它的使用计数器
//被设置为1 (由于 put 操作)。对缓存中的键执行get或put操作，使用计数器的值将会递增。函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。

//示例：
//输入：
//["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get", "get"]
//[[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]]
//输出：
//[null, null, null, 1, null, -1, 3, null, -1, 3, 4]
//解释：
//// cnt(x) = 键 x 的使用计数
//// cache=[] 将显示最后一次使用的顺序（最左边的元素是最近的）
//LFUCache lfu = new LFUCache(2);
//lfu.put(1, 1);   // cache=[1,_], cnt(1)=1
//lfu.put(2, 2);   // cache=[2,1], cnt(2)=1, cnt(1)=1
//lfu.get(1);      // 返回 1 ache=[1,2], cnt(2)=1, cnt(1)=2
//lfu.put(3, 3);   // 去除键 2 ，因为 cnt(2)=1 ，使用计数最小 cache=[3,1], cnt(3)=1, cnt(1)=2
//lfu.get(2);      // 返回 -1（未找到）
//lfu.get(3);      // 返回 3  cache=[3,1], cnt(3)=2, cnt(1)=2
//lfu.put(4, 4);   // 去除键 1 ，1 和 3 的 cnt 相同，但 1 最久未使用 cache=[4,3], cnt(4)=1, cnt(3)=2
//lfu.get(1);      // 返回 -1（未找到）
//lfu.get(3);      // 返回 3 cache=[3,4], cnt(4)=1, cnt(3)=3
//lfu.get(4);      // 返回 4 cache=[3,4], cnt(4)=2, cnt(3)=3

//提示：
//	1 <= capacity <= 10^4
//	0 <= key <= 10^5
//	0 <= value <= 10^9
//	最多调用 2 * 10^5 次 get 和 put 方法

type LFUCache struct {
	capacity int
	keyMap   map[int]*Node
	cntMap   map[int]*LinkedHashMap
	minCnt   int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{capacity, make(map[int]*Node), make(map[int]*LinkedHashMap), 1}
}

func (this *LFUCache) Get(key int) int {
	if node, exist := this.keyMap[key]; exist {
		this.adjust(node)
		return node.val
	} else {
		return -1
	}
}

func (this *LFUCache) Put(key int, value int) {
	if node, exist := this.keyMap[key]; exist {
		node.val = value
		this.adjust(node)
		return
	}
	node := &Node{key, value, 0, nil, nil}
	if len(this.keyMap) == this.capacity {
		dropNode := this.cntMap[this.minCnt].Drop()
		delete(this.keyMap, dropNode.key)
		if this.cntMap[this.minCnt].IsEmpty() {
			delete(this.cntMap, this.minCnt)
		}
	}
	this.adjust(node)
	this.minCnt, this.keyMap[key] = 1, node
}

func (this *LFUCache) adjust(node *Node) {
	if linkNode, ok := this.cntMap[node.cnt]; ok {
		linkNode.Del(node)
		if linkNode.IsEmpty() && this.minCnt == node.cnt {
			this.minCnt++
		}
	}
	node.cnt++
	if _, ok := this.cntMap[node.cnt]; !ok {
		this.cntMap[node.cnt] = LinkConstructor()
	}
	this.cntMap[node.cnt].Add(node)
}

type Node struct {
	key  int
	val  int
	cnt  int
	pre  *Node
	next *Node
}

type LinkedHashMap struct {
	size int
	head *Node
	tail *Node
}

func LinkConstructor() *LinkedHashMap {
	head, tail := &Node{-1, -1, 1, nil, nil}, &Node{-1, -1, 1, nil, nil}
	head.next, tail.pre = tail, head
	return &LinkedHashMap{0, head, tail}
}

func (this *LinkedHashMap) Add(node *Node) {
	node.pre, node.next, this.head.next, this.head.next.pre = this.head, this.head.next, node, node
	this.size++
}

func (this *LinkedHashMap) Del(node *Node) {
	node.pre.next, node.next.pre = node.next, node.pre
	this.size--
}

func (this *LinkedHashMap) IsEmpty() bool {
	return this.size == 0
}

func (this *LinkedHashMap) Drop() *Node {
	node := this.tail.pre
	this.Del(node)
	return node
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
