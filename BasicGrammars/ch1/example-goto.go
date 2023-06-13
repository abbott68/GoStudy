package main

import "fmt"

func main() {
	i := 1

	goto Label

Loop:
	fmt.Println(i)
	i++

	if i <= 5 {
		goto Loop
	}

Label:
	fmt.Println("Loop end")
}
