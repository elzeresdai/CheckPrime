package pkg

import "math"

func CheckPrimes(n []int) []string {

	var result []string
	for _, v := range n {
		result = append(result, isPrime(v))
	}
	return result
}

func isPrime(n int) string {

	if n == 2 {
		return "true"
	}
	if n == 1 {
		return "false"
	}

	num := float64(n)
	sqrt := int(math.Sqrt(num))

	for i := 2; i <= sqrt; i++ {
		if n%i == 0 {
			return "false"
		}
	}

	return "true"
}
