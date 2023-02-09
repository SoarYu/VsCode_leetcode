/*
 * @lc app=leetcode.cn id=1797 lang=golang
 *
 * [1797] 设计一个验证系统
 */
package daily

// @lc code=start
type AuthenticationManager struct {
	timeToLive        int
	tokenExpiredTable map[string]int
}

func Constructor(timeToLive int) AuthenticationManager {
	return AuthenticationManager{timeToLive: timeToLive, tokenExpiredTable: make(map[string]int)}
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
	this.tokenExpiredTable[tokenId] = currentTime + this.timeToLive
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
	if val, hasVal := this.tokenExpiredTable[tokenId]; hasVal && val > currentTime {
		// 更新
		this.tokenExpiredTable[tokenId] = currentTime + this.timeToLive
	} else {
		// 不存在或以过期
	}
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	// 在给定 currentTime 时刻，未过期 的验证码数目 currentTime严格递增
	var tokens int
	for token, expiredTime := range this.tokenExpiredTable {
		if expiredTime <= currentTime { // 已过期
			delete(this.tokenExpiredTable, token)
		} else {
			tokens++
		}
	}
	return tokens
}

/**
 * Your AuthenticationManager object will be instantiated and called as such:
 * obj := Constructor(timeToLive);
 * obj.Generate(tokenId,currentTime);
 * obj.Renew(tokenId,currentTime);
 * param_3 := obj.CountUnexpiredTokens(currentTime);
 */
// @lc code=end
