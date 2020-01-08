package multiplication

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

// Karatsuba multiplication
// TODO: make correct. I think it is failing to handle odd numbers of digits
// TODO: refer to https://en.wikipedia.org/wiki/Karatsuba_algorithm
// TODO: refer to karatsuba in math/big/nat.go
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
	a, _ := new(big.Int).SetString(x[:n/2], 10)
	b, _ := new(big.Int).SetString(x[n/2:], 10)
	c, _ := new(big.Int).SetString(y[:n/2], 10)
	d, _ := new(big.Int).SetString(y[n/2:], 10)

	p := big.NewInt(0).Add(a, b)

	// p := add(a, b)

	// q := add(c, d)
	q := big.NewInt(0).Add(c, d)

	fmt.Printf("%sa:%s + b:%s = p:%s\n", strings.Repeat("\t", indent), a, b, p)
	fmt.Printf("%sc:%s + d:%s = q:%s\n", strings.Repeat("\t", indent), c, d, q)

	ac := multiply(a.String(), c.String(), indent+1)
	fmt.Printf("%sa x c = ", strings.Repeat("\t", indent))
	fmt.Printf("%s\n", ac)

	bd := multiply(b.String(), d.String(), indent+1)
	fmt.Printf("%sb x d = ", strings.Repeat("\t", indent))
	fmt.Printf("%s\n", bd)

	pq := multiply(p.String(), q.String(), indent+1)
	fmt.Printf("%sp x q = ", strings.Repeat("\t", indent))
	fmt.Printf("%s\n", pq)

	fmt.Printf("%spq: %s, ac: %s, bd: %s\n", strings.Repeat("\t", indent), pq, ac, bd)

	fmt.Printf("%sadbc = %s - %s - %s\n", strings.Repeat("\t", indent), pq, ac, bd)

	pqB, _ := new(big.Int).SetString(pq, 10)
	acB, _ := new(big.Int).SetString(ac, 10)
	adbc := big.NewInt(0).Sub(pqB, acB)
	bdB, _ := new(big.Int).SetString(bd, 10)
	adbc.Sub(adbc, bdB)
	// adbc := add(pq, "-"+ac)
	// adbc = add(adbc, "-"+bd)
	// fmt.Printf("%sadbc: %s\n", strings.Repeat("\t", indent), adbc)
	fmt.Printf("%sac:%s * 10**n: %s\n", strings.Repeat("\t", indent), ac, ac+strings.Repeat("0", n))
	fmt.Printf("%sadbc:%s * 10**(n/2): %s\n", strings.Repeat("\t", indent), adbc.String(), adbc.String()+strings.Repeat("0", n/2))

	ac10n, _ := new(big.Int).SetString(acB.String()+strings.Repeat("0", n), 10)
	adbc10nD2, _ := new(big.Int).SetString(adbc.String()+strings.Repeat("0", n/2), 10)

	total := new(big.Int).Add(ac10n, adbc10nD2)
	total.Add(total, bdB)
	// total := add(ac10n, adbc10nD2)
	// total = add(total, bd)
	fmt.Printf("%s%s x %s totals to: %s\n", strings.Repeat("\t", indent), x, y, total.String())
	return total.String()
}

func multiplyRecursive(x, y string, indent int) string {
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
	a, _ := new(big.Int).SetString(x[:n/2], 10)
	b, _ := new(big.Int).SetString(x[n/2:], 10)
	c, _ := new(big.Int).SetString(y[:n/2], 10)
	d, _ := new(big.Int).SetString(y[n/2:], 10)

	p := big.NewInt(0).Add(a, b)

	// p := add(a, b)

	// q := add(c, d)
	q := big.NewInt(0).Add(c, d)

	fmt.Printf("%sa:%s + b:%s = p:%s\n", strings.Repeat("\t", indent), a, b, p)
	fmt.Printf("%sc:%s + d:%s = q:%s\n", strings.Repeat("\t", indent), c, d, q)

	ac := multiplyRecursive(a.String(), c.String(), indent+1)
	fmt.Printf("%sa x c = ", strings.Repeat("\t", indent))
	fmt.Printf("%s\n", ac)

	ad := multiplyRecursive(a.String(), d.String(), indent+1)

	bc := multiplyRecursive(b.String(), c.String(), indent+1)

	bd := multiplyRecursive(b.String(), d.String(), indent+1)
	fmt.Printf("%sb x d = ", strings.Repeat("\t", indent))

	acB, _ := new(big.Int).SetString(ac, 10)
	adB, _ := new(big.Int).SetString(ad, 10)
	bcB, _ := new(big.Int).SetString(bc, 10)
	bdB, _ := new(big.Int).SetString(bd, 10)

	adbc := new(big.Int).Add(adB, bcB)
	fmt.Printf("%sac:%s * 10**n: %s\n", strings.Repeat("\t", indent), ac, ac+strings.Repeat("0", n))
	fmt.Printf("%sad+bc:%s * 10**(n/2): %s\n", strings.Repeat("\t", indent), adbc.String(), adbc.String()+strings.Repeat("0", n/2))

	ac10n, _ := new(big.Int).SetString(acB.String()+strings.Repeat("0", n), 10)
	adbc10nD2, _ := new(big.Int).SetString(adbc.String()+strings.Repeat("0", n/2), 10)

	total := new(big.Int).Add(ac10n, adbc10nD2)
	total.Add(total, bdB)
	// total := add(ac10n.String(), adbc10nD2.String())
	// total = add(total, bdB.String())
	fmt.Printf("%stotal: %s\n", strings.Repeat("\t", indent), total.String())
	return total.String()
	// return total
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

// TODO: make work with values which would produce a negative result
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
		signX := signX
		signY := signY
		a, _ := strconv.Atoi(string(signX + x[i:i+1]))
		b, _ := strconv.Atoi(string(signY + y[i:i+1]))
		// ensure largest is on top
		if a < b {
			a, b = b, a
			signX, signY = signY, signX
		}
		var p int
		if signX == signY {
			p = a + b
		}

		p = (carry + a) + b
		signPs, ps := sign(strconv.Itoa(p))
		if signPs == "-" && (carry+a) > b {
			carry = -1
			result = append([]byte{ps[0]}, result...)
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
	total := string(result)
	signTotal, total := sign(total)
	total = strings.TrimLeft(total, "0")
	if len(total) == 0 {
		return "0"
	}
	if signTotal == "-" {
		total = signTotal + total
	}
	return total
}
