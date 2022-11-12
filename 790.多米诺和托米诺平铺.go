/*
 * @lc app=leetcode.cn id=790 lang=golang
 *
 * [790] 多米诺和托米诺平铺
 */
package main

// @lc code=start
func numTilings(n int) int {
	f := [4]int{}
	f[0] = 1 // f[0][0]=1
	const mod int = 1e9 + 7
	for i := 1; i <= n; i++ { // f[i][0-3]
		g := [4]int{}
		g[0] = (f[0] + f[1] + f[2] + f[3]) % mod
		g[1] = (f[2] + f[3]) % mod
		g[2] = (f[1] + f[3]) % mod
		g[3] = f[0]
		f = g
	}
	return f[0]
}

// @lc code=end

/*
最后一列铺满，记为 0
最后一列只铺了上方一个瓷砖，记为 1
最后一列只铺了下方一个瓷砖，记为 2
最后一列没有铺瓷砖，记为 3

f[n][0] = f[n-1][0] + f[n-1][1]+L + f[n-1][2]+L + f[n-1][3]
f[n][1] = f[n-1][2]+a + f[n-1][3]+L
f[n][2] = f[n-1][1]+a + f[n-1][3]+L
f[n][3] = f[n-1][0]
*/
