package countPrimes

// easy problem

// Count the number of prime numbers less than a non-negative number, n.

func countPrimes(n int) int {
	if n < 2 {
		return 0
	}
	// make sieve
	isPrime := make([]bool, n)
	//  init it with true
	for i := range isPrime {
		isPrime[i] = true
	}
	//  find non-primes by iterating p from 2 to the sqrt of n
	for p := 2; p*p <= n; p++ {
		//  	for each p iterate j from p*p to n by p, marking each j as non-prime
		for j := p * p; j < n; j = j + p {
			isPrime[j] = false
		}
	}
	//  count primes by iterating i from 2 to n incrementing count for each true
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
