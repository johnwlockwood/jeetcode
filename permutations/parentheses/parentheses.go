package parentheses

// refered to solution
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	perms := make([]string, 0)
	helper(n, 0, 0, "", &perms)
	return perms
}

func helper(max, open, close int, cur string, perms *[]string) {
	if len(cur) == max*2 {
		*perms = append(*perms, cur)
		return
	}
	if open < max {
		helper(max, open+1, close, cur+"(", perms)
	}
	if close < open {
		helper(max, open, close+1, cur+")", perms)
	}
}
