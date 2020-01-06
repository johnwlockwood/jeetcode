package largest

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		is := strconv.Itoa(nums[i])
		js := strconv.Itoa(nums[j])
		ij, err := strconv.ParseInt(is+js, 10, 64)
		if err != nil {
			return false
		}
		ji, err := strconv.ParseInt(js+is, 10, 64)
		if err != nil {
			return false
		}
		return ij > ji
	})
	l := ""
	for _, n := range nums {
		l += strconv.FormatInt(int64(n), 10)
	}
	l = strings.TrimLeft(l, "0")
	if len(l) == 0 {
		return "0"
	}
	return l
}

func countDigits(i int) int {
	count := 0
	if i == 0 {
		return 1
	}
	for i != 0 {
		i /= 10
		count++
	}
	return count
}

type byLargestCombo []int

func (a byLargestCombo) Len() int      { return len(a) }
func (a byLargestCombo) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byLargestCombo) Less(i, j int) bool {
	ic := countDigits(a[i])
	jc := countDigits(a[j])

	ij := a[i]*int(math.Pow10(jc)) + a[j]
	ji := a[j]*int(math.Pow10(ic)) + a[i]

	return ij > ji
}

// BenchmarkLargestNumbeB/3,_30,_34,_5,_9-16         	 2068173	       559 ns/op	     184 B/op	       9 allocs/op
// BenchmarkLargestNumbeB/3,_30,_34,_5,_9,_0,_0,_999999997,_999999999-16         	 1329195	       907 ns/op	     288 B/op	      14 allocs/op
// BenchmarkLargestNumbeB/range_3_-_66-16                                        	  190947	      6427 ns/op	    1144 B/op	      58 allocs/op
// math.Pow10 based

// BenchmarkLargestNumbeB/3,_30,_34,_5,_9-16         	 1983177	       591 ns/op	     184 B/op	       9 allocs/op
// BenchmarkLargestNumbeB/3,_30,_34,_5,_9,_0,_0,_999999997,_999999999-16         	 1000000	      1013 ns/op	     288 B/op	      14 allocs/op
// BenchmarkLargestNumbeB/range_3_-_66-16                                        	  138556	      7410 ns/op	    1144 B/op	      58 allocs/op
func largestNumberB(nums []int) string {
	a := byLargestCombo(nums)
	sort.Sort(a)
	b := new(strings.Builder)
	for _, n := range a {
		fmt.Fprint(b, strconv.FormatInt(int64(n), 10))
	}
	l := strings.TrimLeft(b.String(), "0")
	if len(l) == 0 {
		return "0"
	}
	return l
}
