/*
 * @lc app=leetcode.cn id=1781 lang=golang
 *
 * [1781] 所有子字符串美丽值之和
 */
package daily

// @lc code=start
func beautySum(s string) int {
	var ans int
	var dfs func(index int)
	dfs = func(index int) {
		var ch [26]int
		ch[s[index]-'a']++
		ch[s[index+1]-'a']++
		// ch[s[index+2]-'a']++

		for right := index + 2; right < len(s); right++ {
			// fmt.Println(s[index : right+1])
			ch[s[right]-'a']++

			max, min := ch[s[index]-'a'], ch[s[index]-'a']
			for i := range ch {
				if ch[i] != 0 && ch[i] < min {
					min = ch[i]
				}
				if ch[i] > max {
					max = ch[i]
				}
			}
			ans += max - min
		}
	}

	for i := 0; i < len(s)-2; i++ {
		var ch [26]int
		ch[s[i]-'a']++
		ch[s[i+1]-'a']++
		for right := i + 2; right < len(s); right++ {
			// fmt.Println(s[index : right+1])
			ch[s[right]-'a']++

			max, min := ch[s[i]-'a'], ch[s[i]-'a']
			for _, v := range ch {
				if v != 0 && v < min {
					min = v
				}
				if v > max {
					max = v
				}
			}
			ans += max - min
		}
	}

	return ans
}

// @lc code=end
