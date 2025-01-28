package main

import "fmt"

func main() {

	s := "10100100100111101100110111"

	for _, s1 := range s {
		e := string(s1)

		if e == "1" {
			fmt.Print("true, ")
		} else {
			fmt.Print("false, ")
		}
	}

	fmt.Printf("\n\n\n")

	s = "10111000110111000100110111"

	for _, s1 := range s {
		e := string(s1)

		if e == "1" {
			fmt.Print("true, ")
		} else {
			fmt.Print("false, ")
		}
	}

	fmt.Printf("\n\n\n")

	s = "101011101011110110001101110"

	for _, s1 := range s {
		e := string(s1)

		if e == "1" {
			fmt.Print("true, ")
		} else {
			fmt.Print("false, ")
		}
	}

	fmt.Printf("\n\n\n")
}

//  10100100100111101100110111 = 43154231
//  10111000110111000100110111 = 48460087
// 101011101011110110001101110
