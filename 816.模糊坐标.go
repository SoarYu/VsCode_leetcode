/*
 * @lc app=leetcode.cn id=816 lang=golang
 *
 * [816] 模糊坐标
 */
package main

import "strings"

// @lc code=start
func ambiguousCoordinates(s string) []string {
	s = s[1:]
	s = s[:len(s)-1]
	n := len(s)
	res := []string{}
	for i := 0; i < n-1; i++ {
		s1 := s[:i+1]
		s2 := s[i+1 : n]
		arr1, arr2 := []string{}, []string{}
		if isValid(s1) {
			arr1 = append(arr1, s1)
		}
		if isValid(s2) {
			arr2 = append(arr2, s2)
		}
		for x1 := 1; x1 < len(s1); x1++ {
			if tmp := s1[:x1] + "." + s1[x1:]; isValid(tmp) {
				arr1 = append(arr1, tmp)
			}
		}
		for x2 := 1; x2 < len(s2); x2++ {
			if tmp := s2[:x2] + "." + s2[x2:]; isValid(tmp) {
				arr2 = append(arr2, tmp)
			}
		}
		for _, str2 := range arr2 {
			for _, str1 := range arr1 {
				res = append(res, "("+str1+", "+str2+")")
			}
		}
	}
	return res
}

func isValid(s string) bool {
	// 1. 存在前导零
	if len(s) > 1 && s[0] == '0' && s[1] != '.' {
		return false
	}
	// 2. 以零结尾的小数
	if strings.Contains(s, ".") {
		if s[len(s)-1] == '0' {
			return false
		}
	}
	return true
}

// @lc code=end
