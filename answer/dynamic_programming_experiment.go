package answer

import "fmt"

var counterPlainRecursive int = 0
var counterDynamicProgramming int = 0

func DynamicProgrammingExperiment() {
	fmt.Println("\n\nDynamic Programming Experiment")
	fibonacciRecursive(20)
	fibonacciDynamicProgramming(20)

	fmt.Println("We did ", counterPlainRecursive, " calculcations for plain recursive")
	fmt.Println("We did ", counterDynamicProgramming, " calculcations for dynamic programming")
}

func fibonacciRecursive(n int) int {
	counterPlainRecursive++
	if n < 2 {
		return n
	}

	return fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
}

func fibonacciDynamicProgramming(n int) int {
	cache := make(map[int]int)

	var fib func(n int) int
	fib = func(n int) int {
		counterDynamicProgramming++
		cacheFound := cache[n]

		if cacheFound != 0 {
			return cacheFound
		} else {
			if n < 2 {
				return n
			}
			cache[n] = fib(n-1) + fib(n-2)
			return cache[n]
		}
	}

	return fib(n)
}
