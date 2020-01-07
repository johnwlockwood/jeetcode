package multiplication

import (
	"fmt"
	"strconv"
	"strings"
)

func multiply(x, y string, indent int) string {
	if len(x) != len(y) {
		if len(x) < len(y) {
			x = strings.Repeat("0", len(y)-len(x)) + x
		} else {
			y = strings.Repeat("0", len(x)-len(y)) + y
		}
	}
	fmt.Printf("%s%s x %s\n", strings.Repeat("\t", indent), x, y)
	n := len(x)
	if n == 1 {
		a, _ := strconv.Atoi(x)
		b, _ := strconv.Atoi(y)
		total := strconv.Itoa(a * b)
		fmt.Printf("%s= %s\n", strings.Repeat("\t", indent), total)
		return total
	}
	a, b := x[:n/2], x[n/2:]
	c, d := y[:n/2], y[n/2:]
	fmt.Printf("%sa: %s, b: %s, c: %s, d: %s\n", strings.Repeat("\t", indent), a, b, c, d)
	p := add(a, b)
	q := add(c, d)
	fmt.Printf("%sa:%s + b:%s = p: %s\n", strings.Repeat("\t", indent), a, b, p)
	fmt.Printf("%sc:%s + d:%s = q:%s\n", strings.Repeat("\t", indent), c, d, q)
	fmt.Printf("%sa x c\n", strings.Repeat("\t", indent))
	ac := multiply(a, c, indent+1)
	fmt.Printf("%sb x d\n", strings.Repeat("\t", indent))
	bd := multiply(b, d, indent+1)
	fmt.Printf("%sp x q\n", strings.Repeat("\t", indent))
	pq := multiply(p, q, indent+1)

	fmt.Printf("%spq: %s, ac: %s, bd: %s\n", strings.Repeat("\t", indent), pq, ac, bd)
	adbc := add(add(pq, "-"+ac), "-"+bd)
	// fmt.Printf("%sadbc: %s\n", strings.Repeat("\t", indent), adbc)
	fmt.Printf("%sac:%s * 10**n: %s\n", strings.Repeat("\t", indent), ac, ac+strings.Repeat("0", n))
	fmt.Printf("%sadbc:%s * 10**(n/2): %s\n", strings.Repeat("\t", indent), adbc, adbc+strings.Repeat("0", n/2))

	ac10n := ac + strings.Repeat("0", n)
	adbc10nD2 := adbc + strings.Repeat("0", n/2)
	total := add(ac10n, adbc10nD2)
	total = add(total, bd)
	total = strings.TrimLeft(total, "0")
	if len(total) == 0 {
		total = "0"
	}
	fmt.Printf("%stotal: %s\n", strings.Repeat("\t", indent), total)
	return total
}

func sign(x string) (string, string) {
	if len(x) > 1 {
		switch x[0] {
		case '-':
			return "-", x[1:]
		case '+':
			return "+", x[1:]
		}
	}
	return "+", x
}

func add(x, y string) string {
	signX, x := sign(x)
	signY, y := sign(y)

	if len(x) != len(y) {
		if len(x) < len(y) {
			x = strings.Repeat("0", len(y)-len(x)) + x
		} else {
			y = strings.Repeat("0", len(x)-len(y)) + y
		}
	}
	n := len(x)
	carry := 0
	result := make([]byte, 0)
	for i := n - 1; i >= 0; i-- {
		a, _ := strconv.Atoi(string(signX + x[i:i+1]))
		b, _ := strconv.Atoi(string(signY + y[i:i+1]))
		p := a + b + carry
		signPs, ps := sign(strconv.Itoa(p))
		if signPs == "-" {
			carry = -1
		} else if len(ps) > 1 {
			carry, _ = strconv.Atoi(ps[:1])
			result = append([]byte{ps[1]}, result...)
		} else {
			carry = 0
			result = append([]byte{ps[0]}, result...)
		}
	}
	if carry != 0 {
		cs := []byte(strconv.Itoa(carry))
		result = append(cs, result...)
	}
	return string(result)
}
