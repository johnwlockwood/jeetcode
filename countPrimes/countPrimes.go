package countPrimes

// Count the number of prime numbers less than a non-negative number, n.

func countPrimes(n int) int {
	primes := makeSieve(n)
	var count int
	for _, v := range primes {
		if v >= n {
			break
		}
		count++
	}
	return count
}

func makeSieve(n int) []int {
	s := 2
	if n < s {
		return []int{}
	}
	sieve := make([]bool, n-s)
	for i := 0; i < n-s-1; i++ {

		p := i + s
		if sieve[i] {
			continue
		}
		for j := p; j+p < n; j = j + p {
			sieve[p+j-s] = true
		}
	}
	primes := make([]int, 0)
	for i, v := range sieve {
		if !v {
			primes = append(primes, i+s)
		}
	}
	return primes
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
