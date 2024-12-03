package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	//Getting parameters
	number, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Panic(err)
	}

	//Fibonacci using recusive function
	result := fibonacci(number)
	fmt.Printf("Recursive: %v\n", result)

	//Fibonacci using Dynamic Function
	memo := make([]int, number+1, number+1)
	resultDynamic := fibonacciDynamic(number, memo)
	fmt.Printf("Dynamic: %v\n", resultDynamic)
}

func fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func fibonacciDynamic(n int, memo []int) int {
	if n == 0 || n == 1 {
		return n
	}

	if memo[n] == 0 {
		memo[n] = fibonacciDynamic(n-1, memo) + fibonacciDynamic(n-2, memo)
	}

	return memo[n]
}

// func fibonacci3(n int) int {
// 	fibMap := make(map[int]int)
// 	for i := 0; i < n; i++ {
// 		if i < 3 {
// 			fibMap[i] = i
// 		} else {
// 			fibMap[i] = fibMap[i-1] + fibMap[i-2]
// 		}
// 	}
// 	return fibMap[n-1]
// }
