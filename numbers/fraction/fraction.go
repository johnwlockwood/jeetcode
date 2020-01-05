package fraction

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// /problems/fraction-to-recurring-decimal/
func fractionToDecimal(numerator int, denominator int) string {
	fmt.Printf("fraction of %d / %d\n", numerator, denominator)
	if numerator%denominator == 0 {
		fmt.Printf("\t%s\n", strconv.Itoa(numerator/denominator))
		return strconv.Itoa(numerator / denominator)
	}
	fraction := new(strings.Builder)

	dividend := int64(math.Abs(float64(numerator)))
	divisor := int64(math.Abs(float64(denominator)))

	if numerator < 0 && denominator >= 0 || numerator >= 0 && denominator < 0 {
		fmt.Fprint(fraction, "-")
	}
	fmt.Fprint(fraction, strconv.Itoa(int(dividend/divisor)))
	fmt.Fprint(fraction, ".")

	remainder := dividend % divisor

	fmt.Printf("\tfraction %s\n", fraction.String())
	rm := make(map[int64]int)
	for remainder != 0 {
		if l, ok := rm[remainder]; ok {
			f := fraction.String()
			before, last := f[:l], f[l:]
			fraction.Reset()
			fmt.Fprintf(fraction, "%s(%s)", string(before), string(last))
			break
		}
		rm[remainder] = fraction.Len()
		fmt.Printf("\tremainder %d", remainder)
		remainder *= 10
		fmt.Printf(" * 10 = %d; /d: %d : %d\n", remainder, divisor, remainder/divisor)
		fmt.Fprintf(fraction, "%d", remainder/divisor)
		remainder %= divisor
	}
	fmt.Printf("\t%s\n", fraction.String())
	return fraction.String()
}
