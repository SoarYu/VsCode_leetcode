package daily
import(
	"strings"
)
// @lc code=start
func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	return strings.Join(word1, "") == strings.Join(word2, "")
}

func arrayStringsAreEqual0(word1 []string, word2 []string) bool {
	p1, p2 := 0, 0
	w1, w2 := 0, 0
	for w1 < len(word1) && w2 < len(word2) {
		if word1[w1][p1] != word2[w2][p2] {
			return false
		}
		p1, p2 = p1+1, p2+1
		if p1 == len(word1[w1]) {
			w1++
			p1 = 0
		}
		if p2 == len(word2[w2]) {
			w2++
			p2 = 0
		}
	}

	return w1 == len(word1) && w2 == len(word2)
}


// @lc code=end

