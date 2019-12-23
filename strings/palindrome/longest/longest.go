package longest

import "fmt"

func exploreEvenPalindrome(s string, start, end int) bool {
	// accepts start: and :end of known even palindrome aa of [baab] 1:3
	// is the next and prev also matching?
	next := end + 1
	prev := start - 1
	if s[prev] == s[next] {
		return true
	}
	return false
}

func exploreOddPalindrome(s string, start, end int) bool {
	// accepts start: and :end of known odd palindrome bab of [ababa] 1:4
	// is the next and prev also matching?
	next := end + 1
	prev := start - 1
	if s[prev] == s[next] {
		return true
	}
	return false
}

// odd palindrome is bab
// even palindrome is abba
func constructPalindrome(s string, index int) (int, int) {
	// returns start: and :end of palindrome [acbcb] 1:4
	// [acbcbaf]
	fmt.Printf("construct Palindrome from %s at index %d\n", s, index)
	prev := index - 1
	next := index + 1
	if index == 0 || index == len(s) {
		if next < len(s) && s[index] == s[next] {
			return index, next + 1
		}
		fmt.Printf("A.1  %d:%d\n", index, next)
		return index, next
	}
	if next == len(s) {
		fmt.Printf("A.2  %d:%d\n", index, next)
		return index, next
	}
	if s[index] == s[next] {
		// could be an even two letter p
		if s[prev] == s[next] {
			// at least odd three letter p
			// advance prev and next and try more prev and next
			// if not more, then use this prev and next
			i := 0
			for prev-i > 0 && next+i < len(s) {
				// if exploreOddPalindrome(s, prev-i, next+i)
				if !exploreOddPalindrome(s, prev-i, next+i) {
					i-- // backtrack
					fmt.Println("B")
					fmt.Printf("B  %d:%d\n", prev-i, next+i)
					return prev - i, next + i
				} else if prev-i-1 == 0 || next+i+1 == len(s) {
					// at end, it is a still a palindrome
					fmt.Println("C")
					fmt.Printf("C  %d:%d\n", prev-i-1, next+i+2)
					return prev - i - 1, next + i + 2
				}
				i++
			}
			fmt.Println("D")
			fmt.Printf("D i:%d  %d:%d\n", i, prev, next+1)
			return prev, next + 1
		}
		// only an even two letter p
		fmt.Printf("E  %d:%d\n", prev, index+1)
		return prev, index + 1

	} else if s[prev] == s[index] {
		// could be an even two letter p
		// could be a larger even p
		fmt.Println("could be an even two letter p")
		fmt.Println("could be a larger even p")
		i := 1
		for prev-i > 0 && index+i < len(s) {
			if !exploreEvenPalindrome(s, prev-i, index+i) {
				i-- // backtrack
				fmt.Printf("F i:%d  %d:%d\n", i, prev-i, index+i)
				return prev - i, index + i
			} else if prev-i-1 == 0 || index+i+1 == len(s) {
				// at end, it is a still a palindrome
				fmt.Printf("G i:%d  %d:%d\n", i, prev-i-1, index+i+2)
				return prev - i - 1, index + i + 2
			}
			fmt.Printf(".")
			i++
		}
		// only an even two letter p
		fmt.Println("H")
		return prev, index + 1

	} else if s[prev] == s[next] {
		// could be an odd three letter
		i := 0
		for prev-i > 0 && next+i < len(s) {
			// if exploreOddPalindrome(s, prev-i, next+i)
			if !exploreOddPalindrome(s, prev-i, next+i) {
				i-- // backtrack
				fmt.Println("I")
				return prev - i, next + i
			} else if prev-i-1 == 0 || next+i+1 == len(s) {
				// at end, it is a still a palindrome
				fmt.Println("J")
				return prev - i - 1, next + i + 2
			}
			i++
		}
		fmt.Println("K")
		return prev, next + 1
	}
	// only one letter
	fmt.Println("L")
	return index, index + 1
}

func reverse(s string) []rune {
	t := make([]rune, len(s))
	j := len(s) - 1
	for _, v := range s {
		t[j] = v
		j--
	}
	return t
}

// compute palindrome. is it only one letter?
// expand start and end until not a palindrome
//

func isPalindrome(s string, at int) (bool, int, int) {
	fmt.Printf("is pal at %d: %s\n", at, s)
	left, leftGood := at-1, at-1
	right, rightGood := at+2, at+2
	is := false
	for left < right && left >= 0 && right < len(s) {
		if string(s[left:right]) == string(reverse(s[left:right])) {
			is = false
			leftGood = left
			rightGood = right
			fmt.Printf("pal: %s\n", string(s[left:right]))
		} else {
			reverse(s[left:right])
			fmt.Printf(" %v %v %v \n", s[left], s[at], s[right-1])
			if s[left] == s[at] {
				fmt.Println("left 2")
				return true, left, at + 1
			} else if s[at] == s[right-1] {
				fmt.Println("right 2")
				return true, at, right
			}
		}
		if is {
			right++
			left++
		} else {
			right--
			left--
			break
		}
	}

	return is, leftGood, rightGood
}

func longestPalindrome(s string) string {
	// locate and analyze?
	// maybe find
	return ""
}
