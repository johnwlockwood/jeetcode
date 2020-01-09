package words

// Write a function that takes an integer between 0 and 999 (inclusive) and return its English representation.
// Input/output examples:

// function(5) -> “five”
// function(897) -> “eight hundred ninety seven”

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	testNumberWords()
}

func digitWord(n, place int) string {
	ones := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	// fmt.Printf("\tn: %d place %d\n", n, place)
	if n < 10 && n >= 0 && n < len(ones) {
		return ones[n]
	}
	return "wat"
}

func digits(n int) []int {
	if n == 0 {
		return []int{0}
	}
	values := make([]int, 0)
	if n < 0 {
		return values
	}

	for i := 0; int(math.Pow10(i)) <= n; i++ {
		dig := (n % int(math.Pow10(i+1))) / int(math.Pow10(i))
		values = append(values, dig)
	}
	return values
}

func reverse(nums []int, low, high int) {
	if high <= low {
		return
	}
	for l, r := low, high; l < r; l++ {
		nums[l], nums[r] = nums[r], nums[l]
		r--
	}
}

func numberWords(n int) string {
	fmt.Printf("number: %d\n", n)
	ones := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	words := make([]string, 0)
	ds := digits(n)
	reverse(ds, 0, len(ds)-1)
	fmt.Printf("digits %v\n", ds)
DigitLoop:
	for i := 0; i < len(ds); i++ {
		place := len(ds) - i
		switch place {
		case 3:
			words = append(words, ones[ds[i]], "hundred")
			if ds[i+1] == 0 && ds[i+2] == 0 {
				break DigitLoop
			}
		case 2:
			if ds[i] == 0 {
				continue
			}
			switch ds[i] {
			case 0:
				continue
			case 1:
				switch ds[i+1] {
				case 0:
					words = append(words, "ten")
				case 1:
					words = append(words, "eleven")
				case 2:
					words = append(words, "twelve")
				case 3:
					words = append(words, "thirteen")
				case 4:
					words = append(words, "fourteen")
				case 5:
					words = append(words, "fifteen")
				case 6:
					words = append(words, "sixteen")
				case 7:
					words = append(words, "seventeen")
				case 8:
					words = append(words, "eighteen")
				case 9:
					words = append(words, "nineteen")
				}
				break DigitLoop
			case 2:
				words = append(words, "twenty")
			case 3:
				words = append(words, "thirty")
			case 4:
				words = append(words, "fourty")
			case 5:
				words = append(words, "fifty")
			case 6:
				words = append(words, "sixty")
			case 7:
				words = append(words, "seventy")
			case 8:
				words = append(words, "eighty")
			case 9:
				words = append(words, "ninety")
			}
		case 1:
			words = append(words, digitWord(ds[i], place))
		}
	}

	return strings.Join(words, " ")
}

func testNumberWords() {
	if got := numberWords(3); got != "three" {
		fmt.Printf("got %s, want %s\n", got, "three")
	}
	if got := numberWords(4); got != "four" {
		fmt.Printf("got %s, want %s\n", got, "four")
	}
	if got := numberWords(2); got != "two" {
		fmt.Printf("got %s, want %s\n", got, "two")
	}
	if got := numberWords(10); got != "ten" {
		fmt.Printf("got %s, want %s\n", got, "ten")
	}
	if got := numberWords(22); got != "twenty two" {
		fmt.Printf("got %s, want %s\n", got, "twenty two")
	}
	if got := numberWords(0); got != "zero" {
		fmt.Printf("got %s, want %s\n", got, "zero")
	}

	if got := numberWords(100); got != "one hundred" {
		fmt.Printf("got %s, want %s\n", got, "one hundred")
	}
	if got := numberWords(200); got != "two hundred" {
		fmt.Printf("got %s, want %s\n", got, "two hundred")
	}
	if got := numberWords(210); got != "two hundred ten" {
		fmt.Printf("got %s, want %s\n", got, "two hundred ten")
	}
}

// 9
// nine
// two

// 10
// ten
// 12
// ten two
//

// 20
// twenty
// 22
// twenty two

// 200
// two hundred
// 300
// three hundred
// 999
// nine hundred ninety nine
