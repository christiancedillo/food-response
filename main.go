package main

import (
	"fmt"
)

func main() {
	fmt.Println("what you wanna eat boss?")

	var foodOption string
	fmt.Scan(&foodOption)

	fmt.Printf("%v sounds good", foodOption)
}
