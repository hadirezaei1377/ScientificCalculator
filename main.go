package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello!")
	for {
		// get input
		input := getInput()
		// parse input
		n1, n2, op := parse(input)
		// calculate
		result := calculate(n1, n2, op)
		// print result
		fmt.Println(result)

	}
}

func getInput() string {
	var input string
	fmt.Print(">")
	fmt.Scanln(&input)
	return input
}

// our inputs is like 2 + 3
// + is string
// we consider all as string then parse it

// parsing
func parse(input string) (number_1, number_2 float64, operand rune) {
	for _, char := range "+-*/%^" {
		// convert inputs to  a slice(by split)
		// subs = sub string ,A subset of strings ,subs can 1 or 0
		subs := strings.Split(input, string(char))
		if len(subs) != 1 {
			// we have string and numbers want int
			// ignore err as inputs prase in strconv
			number_1, _ := strconv.ParseFloat(subs[0], 64)
			number_2, _ := strconv.ParseFloat(subs[1], 64)
			operand = char
			return number_1, number_2, operand
		}
	}
	return 0, 0, 0
}

// calculate function
func calculate(number_1, number_2 float64, operand rune) float64 {
	switch operand {
	case '+':
		return number_1 + number_2
	case '-':
		return number_1 - number_2
	case '*':
		return number_1 * number_2
	case '/':
		return number_1 / number_2
	case '%':
		return math.Mod(number_1, number_2)
	case '^':
		return math.Pow(number_1, number_2)
	}
	return 0
}
