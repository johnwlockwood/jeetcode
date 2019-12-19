package anagram

func isAnagram(s string, t string) bool {
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
		if ct, ok := countT[k]; ok {
			if cs != ct {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
