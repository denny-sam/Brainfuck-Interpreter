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

	var cellStrip = make([]rune, len(code))
	position := 0
	var result rune

	//looping through the code and modifying cellStrip accordingly
	for index, ch := range code {
		switch string(ch) {
		case "+":
			cellStrip[position]++
		case "-":
			cellStrip[position]++
		case ">":
			position++
		case "<":
			position--
		case ".":
			result += cellStrip[position]
		case ",":
			reader := bufio.NewReader(os.Stdin)
			input, _, err := reader.ReadRune()
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v", err)
			}
			cellStrip[position] = input

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
