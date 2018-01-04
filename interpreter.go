package main

import (
	"bufio"
	"fmt"
	"os"
)

var words = [8]string{">", "<", "+", "-", ".", ",", "[", "]"}
var cellStrip = make([]rune, 30000)

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

//Change the value of cellString according to instruction
func giveOutput(code string) {

	position := 0 //holds position of the current cell
	pointer := 0  //points to the code character being executed

	//looping through the code and modifying cellStrip accordingly
	for ; pointer < len(code); pointer++ {
		ch := code[pointer]
		switch string(ch) {
		case "+":
			cellStrip[position]++
		case "-":
			cellStrip[position]--
		case ">":
			position++
		case "<":
			position--
		case ".":
			fmt.Printf("%s", string(cellStrip[position]))
		case ",":
			reader := bufio.NewReader(os.Stdin)
			input, _, err := reader.ReadRune()
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v", err)
			}
			cellStrip[position] = input
		case "[":
			if cellStrip[position] == 0 {
				for string(code[pointer]) != "]" {
					pointer++
				}
				continue
			} else {
				continue
			}
		case "]":
			if cellStrip[position] == 0 {
				continue
			} else {
				for string(code[pointer]) != "[" {
					pointer--
				}
			}

		}

	}
}

func main() {
	var code string
	var minimised string

	//Input buffer
	input := bufio.NewScanner(os.Stdin)

	//Print welcome message
	fmt.Println("An interpreter for the language Brainfuck \nVersion 1.0")

	fmt.Print("$$$ ")

	//Take the code input until user enters the word 'end'
	for input.Scan() && input.Text() != "end" {
		code += input.Text()
	}
	minimised = minimiseCode(code)
	giveOutput(minimised)

}
