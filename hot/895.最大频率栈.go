/*
 * @lc app=leetcode.cn id=895 lang=golang
 *
 * [895] 最大频率栈
 */

/*
统计频率关键：单调栈
*/
package hot

type FreqNode struct {
	val   int
	freq  int
	index int
	// listIndex int
	next *FreqNode
}

// @lc code=start
type FreqStack struct {
	allSize  int
	headList []*FreqNode
	headMap  map[int]*FreqNode
}

func Constructor895() FreqStack {
	return FreqStack{
		headList: make([]*FreqNode, 0),
		headMap:  make(map[int]*FreqNode),
	}
}

func (fs *FreqStack) Push(val int) {
	newNode := &FreqNode{val: val, freq: 1, index: fs.allSize}
	fs.allSize++

	head, ok := fs.headMap[val]
	if !ok {
		head = &FreqNode{val: val, freq: 1}
		head.index = len(fs.headList)
		fs.headList = append(fs.headList, head)
		fs.headMap[val] = head
	} else {
		head.freq++
		newNode.next = head.next
	}

	head.next = newNode

	// 调整位置
	for i := head.index; i-1 >= 0 && head.freq >= fs.headList[i-1].freq; i-- {
		fs.headList[i-1].index, fs.headList[i].index = i, i-1
		fs.headList[i-1], fs.headList[i] = fs.headList[i], fs.headList[i-1]
	}
}

func (fs *FreqStack) Pop() int {
	// 选择最高的频率同时位置最前的删除
	head := fs.headList[0]

	for i := 1; i < len(fs.headList) && fs.headList[i].freq == head.freq; i++ {
		if fs.headList[i].next.index > head.next.index {
			head = fs.headList[i]
		}
	}

	// fmt.Println(head.val, head.freq)
	head.freq--
	delteNode := head.next
	if head.freq == 0 {
		// delteTop
		// map
		defer delete(fs.headMap, head.val)
	} else {
		head.next = delteNode.next
	}

	// 调整位置
	for i := head.index; i+1 < len(fs.headList) && head.freq <= fs.headList[i+1].freq; i++ {
		fs.headList[i].index, fs.headList[i+1].index = i+1, i
		fs.headList[i], fs.headList[i+1] = fs.headList[i+1], fs.headList[i]
	}

	if head.freq == 0 {
		fs.headList = fs.headList[:len(fs.headList)-1]
	}

	return head.val

}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * param_2 := obj.Pop();
 */
// @lc code=end
