/*
 * @lc app=leetcode.cn id=1604 lang=golang
 *
 * [1604] 警告一小时内使用相同员工卡大于等于三次的人
 */
package daily

import "sort"

// @lc code=start
func alertNames(keyName []string, keyTime []string) []string {
	keyMap := make(map[string][]string)
	for i, name := range keyName {
		if keyMap[name] == nil {
			keyMap[name] = make([]string, 0)
		}
		keyMap[name] = append(keyMap[name], keyTime[i])
	}

	isLessHour := func(before, now string) bool {
		if before[:2] == "23" {
			return now[:2] == "23"
		}
		bytes := []byte(before)
		if bytes[1] == '9' {
			bytes[0]++
			bytes[1] = '0'
		} else {
			bytes[1]++
		}
		return string(bytes) >= now
	}

	ans := make([]string, 0)

	for name, times := range keyMap {
		sort.Strings(times)
		// cnt, left, next := 1, 0, 0

		for i := 2; i < len(times); i++ {
			if !isLessHour(times[i-2], times[i]) {
				ans = append(ans, name)
				break
			}
		}

		// for i := 1; i < len(times); i++ {
		// 	// 两者间隔小于一小时
		// 	if isLessHour(times[left], times[i]) {
		// 		if cnt == 1 {
		// 			cnt = 2
		// 			next = i
		// 		} else if cnt == 2 {
		// 			ans = append(ans, name)
		// 			break
		// 		}
		// 	} else {
		// 		if cnt == 1 {
		// 			left = i
		// 		} else if cnt == 2 {
		// 			left = next
		// 			next = i
		// 		}
		// 	}
		// }
	}

	// for name, times := range keyMap {
	// 	sort.Strings(times)
	// 	cnt, early, second := 1, times[0], ""
	// 	for i := 1; i < len(times); i++ {
	// 		// 两者间隔小于一小时
	// 		if isLessHour(early, times[i]) {
	// 			if cnt == 1 { // 一小时内第二次打卡
	// 				cnt = 2
	// 				second = times[i]
	// 			} else if cnt == 2 { // 一小时内第三次打卡
	// 				ans = append(ans, name)
	// 				break
	// 			}
	// 		} else {
	// 			// 两者间隔大于一小时
	// 			if cnt == 1 { // 第二次大于第一次
	// 				early = times[i]
	// 				cnt = 1
	// 			} else if cnt == 2 { // 第三次大于第一次
	// 				early = second
	// 				second = times[i]
	// 				cnt = 2
	// 			}
	// 		}
	// 	}
	// }
	return ans
}

// @lc code=end
