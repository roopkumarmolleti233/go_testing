package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// parseAndCalculate parses the input and returns the calculation result or an error
func parseAndCalculate(input string) (float64, error) {
	input = strings.TrimSpace(input)
	parts := strings.Fields(input)
	if len(parts) != 3 {
		return 0, errors.New("invalid input format: expected 'number operator number'")
	}

	num1, err1 := strconv.ParseFloat(parts[0], 64)
	operator := parts[1]
	num2, err2 := strconv.ParseFloat(parts[2], 64)

	if err1 != nil || err2 != nil {
		return 0, errors.New("invalid number(s)")
	}

	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, errors.New("division by zero is not allowed")
		}
		return num1 / num2, nil
	default:
		return 0, errors.New("unsupported operator, use +, -, *, or /")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Improved Calculator")
	fmt.Println("-------------------")

	for {
		fmt.Print("Enter expression (e.g., 5 + 3) or 'exit' to quit: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.ToLower(input) == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		result, err := parseAndCalculate(input)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Printf("Result: %.4f\n", result) // show 4 decimal places for precision
	}
}
