/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存
 */
package main

// @lc code=start
type node struct {
	pre  *node
	next *node
	val  int
}

type LRUCache struct {
	capacity int
	curCap   int
	head     *node
	cache    map[int]*node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		head:     &node{},
		cache:    make(map[int]*node),
	}
}

func (this *LRUCache) Get(key int) int {
	n, ok := this.cache[key]
	if ok {
		// 删除n在队中
		n.pre.next = n.next
		n.next.pre = n.pre
		// n放到队尾
		n.pre = this.head.pre
		n.next = this.head
		this.head.pre = n

		if this.head.pre != nil {
			this.head.pre.next = n
		}
		return n.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if this.curCap == this.capacity {
		this.Remove()
	}
	this.curCap++
	n, ok := this.cache[key]
	if ok {
		n.val = value
	} else {
		n = &node{val: value}
		this.cache[key] = n
	}

	// 删除n在队中
	n.pre.next = n.next
	n.next.pre = n.pre
	// n放到队尾
	n.pre = this.head.pre
	n.next = this.head
	this.head.pre = n

}

func (this *LRUCache) Remove() {
	n := this.head.next
	this.head.next = n.next
	n.next.pre = this.head
	delete(this.cache, n.val)
	this.curCap--
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
