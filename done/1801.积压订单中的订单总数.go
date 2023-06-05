/*
 * @lc app=leetcode.cn id=1801 lang=golang
 *
 * [1801] 积压订单中的订单总数
 */
package daily

// @lc code=start
type Order struct {
	price  int
	amount int
	next   *Order
}

/*
buy 最大 先清除小的 递减
sell 最小 先清除大的 递增
*/
func getNumberOfBacklogOrders(orders [][]int) int {
	// 队头
	bHead, sHead := &Order{next: nil}, &Order{next: nil}
	for _, order := range orders {
		newOrder := &Order{price: order[0], amount: order[1], next: nil}
		// buy
		if order[2] == 0 {
			// 检查 sell 最小值 sell 队头
			sNode := sHead
			for sNode.next != nil && sNode.next.price <= newOrder.price && newOrder.amount > 0 {
				tmp := sNode.next
				if newOrder.amount < tmp.amount {
					tmp.amount -= newOrder.amount
					newOrder.amount = 0
					break
				} else {
					newOrder.amount -= tmp.amount
					// 删除 tmp
					sNode.next = tmp.next
				}
				sNode = sNode.next
			}

			// 有剩余插入 buy
			if newOrder.amount > 0 {
				bNode := bHead
				if bNode.next == nil {
					bNode.next = newOrder
				} else {
					for bNode.next != nil {
						tmp := bNode.next
						if newOrder.price < tmp.price {
							bNode = bNode.next
						} else if newOrder.price == tmp.price {
							tmp.amount += newOrder.amount
							break
						} else { // newOrder.price > tmp.price
							newOrder.next = tmp
							bNode.next = newOrder
							break
						}
					}
				}
			}
		} else { //sell
			// 检查 buy 最大值 buy 队头
			bNode := bHead
			for bNode.next != nil && bNode.next.price >= newOrder.price && newOrder.amount > 0 {
				tmp := bNode.next
				if newOrder.amount < tmp.amount {
					tmp.amount -= newOrder.amount
					newOrder.amount = 0
					break
				} else {
					newOrder.amount -= tmp.amount
					// 删除 tmp
					bNode.next = tmp.next
				}
				bNode = bNode.next
			}
			// 有剩余插入 sell
			if newOrder.amount > 0 {
				sNode := sHead
				if sNode.next == nil {
					sNode.next = newOrder
				} else {
					for sNode.next != nil {
						tmp := sNode.next
						if newOrder.price > tmp.price {
							sNode = sNode.next
						} else if newOrder.price == tmp.price {
							tmp.amount += newOrder.amount
							break
						} else { // newOrder.price > tmp.price
							newOrder.next = tmp
							sNode.next = newOrder
							break
						}
					}
				}
			}
		}
	}
	var ans int
	for bNode := bHead.next; bNode != nil; bNode = bNode.next {
		ans += bNode.amount
	}
	for sNode := sHead.next; sNode != nil; sNode = sNode.next {
		ans += sNode.amount
	}
	return ans
}

// @lc code=end
