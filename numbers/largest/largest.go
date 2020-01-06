package largest

import (
	"fmt"
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

func largestNumber2(nums []int) string {
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
	b := new(strings.Builder)
	for _, n := range nums {
		fmt.Fprint(b, strconv.FormatInt(int64(n), 10))
	}
	l := strings.TrimLeft(b.String(), "0")
	if len(l) == 0 {
		return "0"
	}
	return l
}
