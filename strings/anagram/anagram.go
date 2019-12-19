package anagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	counter := make([]int, 26)
	for i := 0; i < len(s); i++ {
		counter[s[i]-'a']++
		counter[t[i]-'a']--
	}
	for _, count := range counter {
		if count != 0 {
			return false
		}
	}
	return true
}

// naive approach seems like many have given similar answer
func isAnagramNaive(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	// count each rune
	countS := make(map[rune]int, len(s))
	for _, v := range s {
		if _, ok := countS[v]; ok {
			countS[v]++
		} else {
			countS[v] = 1
		}
	}
	countT := make(map[rune]int, len(t))
	for _, v := range t {
		if _, ok := countT[v]; ok {
			countT[v]++
		} else {
			countT[v] = 1
		}
	}
	for k, cs := range countS {
		ct, ok := countT[k]
		if !ok || ct != cs {
			return false
		}
	}
	return true
}
