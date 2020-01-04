package reverse

import "strings"

func reverseWords(s string) string {
	words := strings.Fields(s)
	reverse(words, 0, len(words)-1)
	return strings.Join(words, " ")
}

func reverse(words []string, low, high int) {
	if high <= low {
		return
	}

	for l, r := low, high; l < r; l++ {
		words[l], words[r] = words[r], words[l]
		r--
	}
}
