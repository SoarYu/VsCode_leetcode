/*
 * @lc app=leetcode.cn id=1754 lang=golang
 *
 * [1754] 构造字典序最大的合并字符串
 */
package daily

// @lc code=start
func largestMerge(word1 string, word2 string) string {
	ans := ""
	i, j := 0, 0
	/*
		1. 字符相同 ，首字符continue，
		2. 非首字符相同， 比较前后，若前面大，break; 其余continue
	*/
	for i < len(word1) && j < len(word2) {
		t1, t2 := i, j
		if word1[t1] == word2[t2] {
			t1, t2 = t1+1, t2+1
			for t1 < len(word1) && t2 < len(word2) && word1[t1] == word2[t2] {
				if word1[t1] < word1[t1-1] || word2[t2] < word2[t2-1] {
					t1--
					t2--
					break
				}
				// bcab
				// bcab
				t1, t2 = t1+1, t2+1
			}
		}
		if word1[t1] >= word2[t2] {
			ans += word1[i : t1+1]
			i = t1 + 1
		} else {
			ans += word2[j : t2+1]
			j = t2 + 1
		}
	}
	if i < len(word1) {
		ans += word1[i:]
	}
	if j < len(word2) {
		ans += word2[j:]
	}
	return ans
}

// @lc code=end



/*

bacabc

badcaba

*/