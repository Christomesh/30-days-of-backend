package day3

import "fmt"

func fibonacci(N int) {
	var fibList []int
	a, b := 0, 1
	fibList = append(fibList, a, b)
	for i := 0; i < N-2; i++ {
		res := a + b
		a = b
		b = res
		fibList = append(fibList, res)
	}
	fmt.Println(fibList)
}
