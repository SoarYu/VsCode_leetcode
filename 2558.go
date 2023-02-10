package daily

func oddString2558(words []string) string {
	n := len(words[0])
	d1 := make([]int, n-1)
	for i := 1; i < n; i++ {
		d1[i-1] = int(words[0][i] - words[0][i-1])
	}
	var count, index [2]int
	index[0] = 0
	count[0] = 1
	for i := 1; i < len(words); i++ {
		word := words[i]
		k := 0
		for j := 1; j < n; j++ {
			d := int(word[j] - word[j-1])
			if k == 0 && d == d1[j-1] {
				continue
			} else {
				k = 1
				break
			}
		}
		index[k] = i
		count[k]++

		if (count[0] > 1 || count[1] > 1) && (count[0] == 1 || count[1] == 1) {
			break
		}
	}

	if count[1] == 1 {
		return words[index[1]]
	} else {
		return words[index[0]]
	}

}
