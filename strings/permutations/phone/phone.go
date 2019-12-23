package phone

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Letter Combinations of a Phone Number.
// Memory Usage: 2.5 MB, less than 54.55% of Go online submissions for Letter Combinations of a Phone Number.
// https://leetcode.com/submissions/detail/288038011/
func letterCombinations(digits string) []string {
	if len(digits) < 1 {
		return []string{}
	}
	digitToLetters := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	lastDigit := digits[len(digits)-1]
	digitLetters := digitToLetters[lastDigit]
	lastCombos := make([]string, 0, len(digitLetters))

	for _, d := range digitLetters {
		lastCombos = append(lastCombos, string(d))
	}
	if len(digits) == 1 {
		return lastCombos
	}
	for i := len(digits) - 2; i >= 0; i-- {
		d := digits[i]
		letters := digitToLetters[d]
		combos := make([]string, 0)
		for _, prefix := range letters {
			for _, suffix := range lastCombos {
				combos = append(combos, string(prefix)+suffix)
			}
		}
		lastCombos = combos
	}

	return lastCombos
}
