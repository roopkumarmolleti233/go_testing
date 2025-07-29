package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Simple Calculator")
    fmt.Println("-----------------")

    for {
        fmt.Print("Enter expression (e.g., 5 + 3) or 'exit' to quit: ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)

        if input == "exit" {
            fmt.Println("Goodbye!")
            break
        }

        parts := strings.Fields(input)
        if len(parts) != 3 {
            fmt.Println("Invalid input format. Please enter like: 5 + 3")
            continue
        }

        num1, err1 := strconv.ParseFloat(parts[0], 64)
        operator := parts[1]
        num2, err2 := strconv.ParseFloat(parts[2], 64)

        if err1 != nil || err2 != nil {
            fmt.Println("Invalid numbers. Please try again.")
            continue
        }

        var result float64
        switch operator {
        case "+":
            result = num1 + num2
        case "-":
            result = num1 - num2
        case "*":
            result = num1 * num2
        case "/":
            if num2 == 0 {
                fmt.Println("Cannot divide by zero.")
                continue
            }
            result = num1 / num2
        default:
            fmt.Println("Unsupported operator. Use +, -, *, or /.")
            continue
        }

        fmt.Printf("Result: %.2f\n", result)
    }
}
