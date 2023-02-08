/*
 * @lc app=leetcode.cn id=1233 lang=golang
 *
 * [1233] 删除子文件夹
 */
package daily

import "strings"

// @lc code=start

func removeSubfolders(folder []string) []string {
	// 字典树
	type trie struct {
		child map[string]*trie
		flag  bool
		val   byte
	}

	root := &trie{}
	ans := make([]string, 0)
	for i := range folder {
		node := root
		for _, f := range strings.Split(folder[i][1:], "/") {
			if node.child == nil {
				node.child = make(map[string]*trie)
			}
			if _, hasT := node.child[f]; !hasT {
				node.child[f] = &trie{}
			}

			node = node.child[f]
			if node.flag {	// 当前文件夹前缀存在，终止循环
				break
			}
		}
		node.flag = true
	}

	for i := range folder {
		node, ok := root, true
		for _, f := range strings.Split(folder[i][1:], "/") {
			if node == nil {
				ok = false
				break
			}
			node = node.child[f]
			if node.flag { // 当前文件夹前缀存在, 让node为nil，再往下就是子文件夹
				node = nil
			}
		}
		if ok {
			ans = append(ans, folder[i])
		}
	}
	return ans
}

// @lc code=end
