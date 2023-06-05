/*
 * @lc app=leetcode.cn id=1268 lang=golang
 *
 * [1268] 搜索推荐系统
 */
package daily

import "sort"

type trie1268 struct {
	children map[byte]*trie
	val      byte
	isEnd    bool
}

func NewTrie() *trie1268 {
	return &trie{children: make(map[byte]*trie)}
}

func (t *trie) insert(word string) {

	tmp := t

	for i := range word {
		if child, hasChild := tmp.children[word[i]]; hasChild {
			tmp = child
		} else {
			newT := NewTrie()
			newT.val = word[i]
			tmp.children[word[i]] = newT
			tmp = newT
		}

		if i == len(word)-1 {
			tmp.isEnd = true
		}
	}

}

func (t *trie) searchAll(s string) []string {
	tmp := t

	res := []string{}

	for i := range s {
		if child, hasChild := tmp.children[s[i]]; !hasChild {
			return res
		} else {
			tmp = child
		}
	}

	// tmp 寻找匹配前缀的词
	var dfs func(*trie, string)
	dfs = func(tt *trie, ss string) {
		if len(res) >= 3 {
			return
		}
		if tt.isEnd {
			res = append(res, ss)
		}
		for i := byte('a'); i <= 'z'; i++ {
			if child, hasChild := tt.children[i]; hasChild {
				dfs(child, ss+string(i))
			}
		}
	}

	dfs(tmp, s)

	return res
}

// @lc code=start
func suggestedProducts(products []string, searchWord string) [][]string {
	sort.Strings(products)

	t := NewTrie()

	for _, word := range products {
		t.insert(word)
	}

	res := make([][]string, len(searchWord))
	for i := range searchWord {
		res[i] = t.searchAll(searchWord[:i+1])
	}
	return res
}

// @lc code=end
