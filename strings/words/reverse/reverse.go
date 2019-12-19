package reverse

func reverse(s string) []rune {
	t := make([]rune, len(s))
	j := len(s) - 1
	for _, v := range s {
		t[j] = v
		j--
	}
	return t
}

func reverseWords(s string) string {
	t := make([]rune, len(s))
	lastWordStart := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			t[i] = ' '
			r := reverse(s[lastWordStart:i])
			for j := 0; j < len(r); j++ {
				t[j+lastWordStart] = r[j]
			}
			lastWordStart = i + 1
		} else if i == len(s)-1 {
			r := reverse(s[lastWordStart : i+1])
			for j := 0; j < len(r); j++ {
				t[lastWordStart+j] = r[j]
			}
		}
	}
	return string(t)
}
