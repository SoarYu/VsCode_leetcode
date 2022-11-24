package daily
// @lc code=start
func halvesAreAlike(s string) bool {
	runeMap := map[byte]bool{'a':true, 'e':true, 'i':true, 'o':true, 'u':true,
	'A':true, 'E':true, 'I':true, 'O':true, 'U':true}
	n := len(s)
	n1, n2 := 0, 0 
	for i, j:=0, n/2; i<n/2; i, j = i+1, j+1 {
		if runeMap[s[i]] {
			n1++
		}
		if runeMap[s[j]] {
			n2++
		}
	} 
	return n1 == n2
}
// @lc code=end

