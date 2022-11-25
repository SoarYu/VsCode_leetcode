package hot

// @lc code=start
type node struct {
	pre  *node
	next *node
	val  int
}

type LRUCache struct {
	capacity int
	size     int
	cache    map[int]*node
	head     *node
	tail     *node
}

func Constructor(capacity int) LRUCache {
	head, tail := &node{}, &node{}
	head.next = tail
	tail.pre = head
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*node),
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	if this.size <= 0 {
		return -1
	}
	n, ok := this.cache[key]
	if ok {
		this.removeNode(n)
		this.addToTail(n)
		return n.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if this.size == this.capacity {
		n := this.removeNode(this.head.next)
		delete(this.cache, n.val)
		this.size--
	}
	this.size++
	n, ok := this.cache[key]
	if ok {
		n.val = value
	} else {
		n = &node{val: value}
		this.cache[key] = n
	}
	this.addToTail(n)
}

func (this *LRUCache) addToTail(n *node) {
	n.pre = this.tail.pre
	n.next = this.tail
	this.tail.pre.next = n
	this.tail.pre = n
}

func (this *LRUCache) removeNode(n *node) *node {
	n.pre.next = n.next
	n.next.pre = n.pre
	return n
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
