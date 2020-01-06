package largest

import (
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
