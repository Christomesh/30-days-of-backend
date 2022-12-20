package day2

import "fmt"

func LeapYear(Y int, N int) {
	var lpy []int
	count := 0
	year := Y

	for count < N {
		if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
			lpy = append(lpy, year)
			count++
		}
		year++
	}
	fmt.Println(lpy)
}
