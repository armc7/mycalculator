package mycalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct{}

func (calc) operate(operation string, operator string) int {
	inputs := strings.Split(operation, operator)

	input1 := parse(inputs[0])
	input2 := parse(inputs[1])

	switch operator {
	case "+":
		return input1 + input2
	case "-":
		return input1 - input2
	case "*":
		return input1 * input2
	case "/":
		return input1 / input2
	default:
		fmt.Println("Invalid input")
		return 0
	}
}

func parse(input string) int {
	result, _ := strconv.Atoi(input)

	return result
}

func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	return scanner.Text()
}
