package parentheses

import (
	"fmt"
	"strings"
)

// refered to solution
func generateParenthesis(n int) []string {
	fmt.Printf("n = %d\n", n)
	if n == 0 {
		return []string{}
	}
	perms := make([]string, 0)
	helper(n, 0, 0, "", &perms)
	return perms
}

func helper(max, open, close int, cur string, perms *[]string) {
	if len(cur) == max*2 {
		fmt.Printf("%s%s\n", strings.Repeat("\t", open+close), cur)
		*perms = append(*perms, cur)
		return
	}
	if open < max {
		fmt.Printf("%s%s\n", strings.Repeat("\t", open+close), cur)
		helper(max, open+1, close, cur+"(", perms)
	}
	if close < open {
		fmt.Printf("%s%s\n", strings.Repeat("\t", open+close), cur)
		helper(max, open, close+1, cur+")", perms)
	}
}
