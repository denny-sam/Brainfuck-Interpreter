package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var code string

	//Print welcome message
	fmt.Println("An interpreter for the language Brainfuck \nVersion 1.0")
	fmt.Println(">>>")

	//Input buffer
	input := bufio.NewScanner(os.Stdin)

	//Take the code input until user enters the word 'end'
	for input.Scan() && input.Text() != "end" {
		code += input.Text()
	}

	//eliminate spaces in the code to get a single line string of instructions
	for _, ch := range code {
		if string(ch) == " " {

		}
	}

}
