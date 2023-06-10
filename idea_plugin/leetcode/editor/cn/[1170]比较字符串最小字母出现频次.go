//定义一个函数 f(s)，统计 s 中（按字典序比较）最小字母的出现频次 ，其中 s 是一个非空字符串。 
// 例如，若 s = "dcce"，那么 f(s) = 2，因为字典序最小字母是 "c"，它出现了 2 次。
//
// 现在，给你两个字符串数组待查表 queries 和词汇表 words 。对于每次查询 queries[i] ，需统计 words 中满足 f(
//queries[i]) < f(W) 的 词的数目 ，W 表示词汇表 words 中的每个词。 
// 请你返回一个整数数组 answer 作为答案，其中每个 answer[i] 是第 i 次查询的结果。
//
// 示例 1： 
//输入：queries = ["cbd"], words = ["zaaaz"]
//输出：[1]
//解释：查询 f("cbd") = 1，而 f("zaaaz") = 3 所以 f("cbd") < f("zaaaz")。
//
// 示例 2： 
//输入：queries = ["bbb","cc"], words = ["a","aa","aaa","aaaa"]
//输出：[1,2]
//解释：第一个查询 f("bbb") < f("aaaa")，第二个查询 f("aaa") 和 f("aaaa") 都 > f("cc")。
// 提示：
// 1 <= queries.length <= 2000
// 1 <= words.length <= 2000 
// 1 <= queries[i].length, words[i].length <= 10 
// queries[i][j]、words[i][j] 都由小写英文字母组成 
// Related Topics 数组 哈希表 字符串 二分查找 排序 👍 87 👎 0
package main
import (
	"sort"
	"fmt"
)
//leetcode submit region begin(Prohibit modification and deletion)
func numSmallerByFrequency(queries []string, words []string) []int {
	m, n := len(queries), len(words)
	res := make([]int, m)
	// 对words字符串的最小字典序数量进行排序
	fwords := make([]int, n)
	for i := range words {
		fwords[i] = fn(words[i])
	}
	sort.Ints(fwords)
	// fmt.Println(fwords)
	// 二分查找有多少个 f(word) > f(query)
	for i := range queries {
		queryCnt := fn(queries[i])
		/**
			该函数使用二分查找的方法，会从[0, n)中取出一个值index
			index为[0, n)中最小的使函数f(index)为True的值，并且f(index+1)也为True
			如果无法找到该index值，则该方法为返回n
		 */
		index := sort.Search(n, func(i int) bool {
			return fwords[i] > queryCnt
		})

		res[i] = n - index

		// fmt.Printf("query:%s\tcnt:%d\tindex:%d\n", queries[i], queryCnt, index)
	}
	return res
}

/**
	计算字符串的最小字典序的字符数量
 */
func fn(s string) int {
	var cnt int
	var minCh byte = 'z'
	for i := range s {
		if s[i] < minCh {
			minCh = s[i]
			cnt = 1
		} else if s[i] == minCh {
			cnt++
		}
	}
	return cnt
}
//leetcode submit region end(Prohibit modification and deletion)

// ["aaabbb","aab","babbab","babbbb","b","bbbbbbbbab","a","bbbbbbbbbb","baaabbaab","aa"]
//     3       2      2         1     1        1        1       10            5      2
//     1 1 1 1 2 2 2 3 5 10
//