package happyNumber

import (
	"math"
)

// IsHappyApproach2 detects a happy number and uses Floyd's Cycle-Finding algorithm to find a cycle
func IsHappyApproach2(n int) bool {
	// Floyd's cycle-finding
	slow := n
	fast := getNext(n)
	for fast != 1 && slow != fast {
		slow = getNext(slow)
		fast = getNext(getNext(fast))
	}
	return fast == 1
}

// IsHappyApproach1 tells you if the number is happy or not
func IsHappyApproach1(n int) bool {
	// a couple of examples to get started
	// n = 7; 7**2 = 49; 4**2 + 9**2; 16+81 = 97; 81+49=130; 1+9+0=10; 1
	// n = 116
	seen := make(map[int]struct{})
	for n != 1 {
		if _, ok := seen[n]; ok {
			break
		}
		if n < 243 { // a number 243 or higher will never get back to itself
			seen[n] = struct{}{}
		}
		n = getNext(n)

	}
	return n == 1
}

// given a number, what is the next number
func getNext(n int) int {
	var totalSum int
	for n > 0 {
		d := n % 10
		n = n / 10
		totalSum = totalSum + d*d
	}
	return totalSum
}

// original naive attempt
func IsHappyNaive(n int) bool {
	// The nth digit is (the remainder of dividing by 10n) divided by 10n-1
	if n < 0 {
		return false
	}
	seen := make(map[int]struct{})
	for n != 1 {
		sum := 0
		theDigits := digits(n)
		// fmt.Println("digits of ", n, " are ", theDigits)
		squares := make([]int, 0, len(theDigits))
		for _, v := range theDigits {
			sq := int(math.Pow(float64(v), float64(2)))
			squares = append(squares, sq)
			sum = sum + sq
		}
		// fmt.Println("sum of squares of ", theDigits, " is ", sum, " squares: ", squares)
		if _, ok := seen[sum]; ok {
			break
		}
		seen[sum] = struct{}{}
		n = sum
	}
	return n == 1
}

func digits(n int) []int {
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
