//
// EPITECH PROJECT, 2020
// 202unsold_2019
// File description:
// main
//

package main

import (
	"fmt"
	"os"
	"strconv"

	functions "./functions"
)

func help() {
	fmt.Printf("USAGE\n   ./205IQ u s [IQ1] [IQ2]\n")
	fmt.Printf("\nDESCRIPTION\n")
	fmt.Printf("   u\t    mean\n")
	fmt.Printf("   s\t    standard deviation\n")
	fmt.Printf("   IQ1\t    minimum IQ\n")
	fmt.Printf("   IQ2\t    maximum IQ\n")
	os.Exit(0)
}

func main() {
	args := os.Args

	if len(args) == 2 {
		if args[1] == "-h" || args[1] == "--help" {
			help()
		}
	}
	if _, err := functions.ErrorArgs(args); err != nil {
		fmt.Fprintf(os.Stderr, "\033[31mX\033[0m Error: %s\n", err)
		os.Exit(84)
	}
	numbers := make([]int, 0)
	for i := 1; i != len(args); i++ {
		nb, err := strconv.Atoi(args[i])
		if err != nil {
			fmt.Println(err, nb)
			os.Exit(84)
		}
		if nb <= 0 || nb > 200 {
			os.Exit(84)
		}
		numbers = append(numbers, nb)
	}
	os.Exit(functions.MathParse(numbers))
}
