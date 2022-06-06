package main

import (
	"fmt"
)

func main() {
	bills := []int{10000, 500, 200, 100, 69, 50, 10, 5, 2, 1}
	var usrplay int
	fmt.Println("How much money you got?")
	fmt.Scanln(&usrplay)
	var i int
	i = 0
	var nbbills int
	fmt.Println("That's")
	for usrplay > 0 {
		nbbills = (usrplay / bills[i])
		if nbbills != 0 {
			fmt.Println(nbbills, " $", bills[i], " bills")
		}
		usrplay = usrplay - (nbbills * bills[i])
		i++

	}
}
