package daily

import "sort"

// @lc code=start
// sort.Slice 自定义排序
func customSortString(order string, s string) string {
	// order字符唯一 s不唯一
	// 顺序
	chMap := make(map[byte]int)
	for i, _ := range order {
		chMap[order[i]] = i
	}
	bSlice := []byte(s)
	sort.Slice(bSlice, func(i, j int) bool {
		return chMap[bSlice[i]] > chMap[bSlice[j]]
	})
	return string(bSlice)
}

// @lc code=end
// 计数排序
func customSortString1(order string, s string) string {
	chMap := make(map[byte]int)
	for i, _ := range s {
		chMap[s[i]]++
	}
	ans := []byte{} // 切片
	// ans := new([]byte)	// 指针
	for i, _ := range order {
		for chMap[order[i]] > 0 {
			ans = append(ans, order[i])
			chMap[order[i]]--
		}
	}
	for k, v := range chMap {
		for v > 0 {
			ans = append(ans, k)
			v--
		}
	}
	return string(ans)
}
