package anagram

import "fmt"

func isAnagram(s string, t string) bool {
	fmt.Println(s, t)
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
	fmt.Printf("\tcountS: %v, countT: %v\n", countS, countT)
	for k, cs := range countS {
		if ct, ok := countT[k]; ok {
			fmt.Printf("for %v s: %d, t: %d\n", string(k), cs, ct)
			if cs != ct {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
