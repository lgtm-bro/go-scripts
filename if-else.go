package main 

import "fmt"

func main() {
	if 7 % 2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	if 8 % 4 == 0 {
		fmt.Println("disvisible by 4")
	}

	if 8 % 2 == 0 || 7 % 2 == 0 {
		fmt.Println("one of these is even")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has one digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}