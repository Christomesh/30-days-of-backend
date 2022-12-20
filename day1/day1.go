package day1

import (
	"fmt"
	"strings"
)

func Dfunc(name string) {
	ch := "! \"#$%&'()*+,-./:;<=>?@[|]\\^_`{|}~"
	for _, v := range name {
		if strings.Contains(ch, string(v)) {
			fmt.Println("Please remove special character or space")
			return
		}
	}
	if strings.ToLower(name) == "bolu" {
		fmt.Println("Welcome Bolu")
		return
	} else if strings.ToLower(name) == "odun" {
		fmt.Println("Welcome Odun")
		return
	}
	fmt.Println("It is nice to meet you")

}
