package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please give me one or more floats.")
		os.Exit(1)
	}
	arguments := os.Args
	min, _ := strconv.ParseFloat(arguments[1], 64)
	max, _ := strconv.ParseFloat(arguments[1], 64)

	for i := 2; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)

		if err != nil {
			fmt.Println("Error Occurred!")
			io.WriteString(os.Stderr, err.Error()+"\n")
			os.Exit(1)
		}

		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}
