package calc

// BinomialCoefficient erwartet zwei Zahlen n und k und liefert den
// Binomialkoeffizienten "n über k".
func BinomialCoefficient(n, k int) int {
	if k == 0 || k == n {
		return 1
	}

	return BinomialCoefficient(n-1, k-1) + BinomialCoefficient(n-1, k)

	//return factorial(n) / (factorial(k) * factorial(n-k))
}

// func factorial(n int) int {
// 	if n == 0 {
// 		return 1
// 	}

// 	return n * factorial(n-1)
// }
