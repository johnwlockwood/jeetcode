package countPrimes

// Count the number of prime numbers less than a non-negative number, n.

func countPrimes(n int) int {
	if n < 2 {
		return 0
	}
	// init primes sieve
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}
	// from 2 to sqrt of n
	for p := 2; p*p <= n; p++ {
		if !isPrime[p] {
			continue
		}
		// for 2p to n increment by p, mark as not prime
		for j := p * p; j < n; j = j + p {
			isPrime[j] = false
		}
	}
	var count int
	for i := 2; i < n; i++ {
		if isPrime[i] {
			count++
		}
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
			// fmt.Printf("\tp %v * j %v == %v, mark sieve[%d]\n", p, j, p+j, p+j-s)
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
