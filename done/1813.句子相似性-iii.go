/*
 * @lc app=leetcode.cn id=1813 lang=golang
 *
 * [1813] 句子相似性 III
 */
package daily

import "strings"

// @lc code=start
func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	/*
	   1. 两侧插入 短插入长的左或右
	   2. 中间插入 短插入长的中间
	*/

	// 2. 中间(以空格为单位), 确认分界点 my name is haley  |   my haley
	var larr, sarr []string
	if arr1, arr2 := strings.Split(sentence1, " "), strings.Split(sentence2, " "); len(arr1) >= len(arr2) {
		larr, sarr = arr1, arr2
	} else {
		larr, sarr = arr2, arr1
	}

	isSame := func(arr1, arr2 []string) bool {
		// fmt.Println(arr1, arr2)
		if len(arr1) != len(arr2) {
			return false
		}
		for i := 0; i < len(arr1); i++ {
			if arr1[i] != arr2[i] {
				return false
			}
		}
		return true
	}

	if isSame(larr[:len(sarr)], sarr) || isSame(larr[len(larr)-len(sarr):], sarr) {
		return true
	}
	// 以第0个词为分界点 到最后
	for i := 1; i < len(sarr); i++ {
		if isSame(larr[:i], sarr[:i]) && isSame(larr[len(larr)-len(sarr)+i:], sarr[i:]) {
			return true
		}
	}
	return false
}

// @lc code=end
