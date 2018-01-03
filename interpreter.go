package main

import (
	"bufio"
	"fmt"
	"os"
)

var words = [8]string{">", "<", "+", "-", ".", ",", "[", "]"}

//eliminate characters not supported to get a single line string of instructions
func minimiseCode(code string) string {
	var minimised string
	for _, ch := range code {
		for _, word := range words {
			if string(ch) == word {
				minimised += string(ch)
			}
		}
	}
	return minimised
}

//Return the result of the program
func giveOutput(code string) string {

	for _, ch := range code {
		switch string(ch) {
		case "+":
			fmt.Println(ch)
		}
	}
	return code
}

func main() {
	var code string
	var minimised string
	var output string

	//Print welcome message
	fmt.Println("An interpreter for the language Brainfuck \nVersion 1.0")
	fmt.Print(">>>")

	//Input buffer
	input := bufio.NewScanner(os.Stdin)

	//Take the code input until user enters the word 'end'
	for input.Scan() && input.Text() != "end" {
		code += input.Text()
	}

	minimised = minimiseCode(code)
	output = giveOutput(minimised)
	fmt.Println(output)

}
