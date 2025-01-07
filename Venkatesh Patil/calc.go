package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    if len(os.Args) < 4 {
        fmt.Println("Usage: <num1> <operator> <num2>")
        return
    }

    num1, _ := strconv.ParseFloat(os.Args[1], 64)
    operator := os.Args[2]
    num2, _ := strconv.ParseFloat(os.Args[3], 64)

    switch operator {
    case "+":
        fmt.Println("Result:", num1+num2)
    case "-":
        fmt.Println("Result:", num1-num2)
    case "*":
        fmt.Println("Result:", num1*num2)
    case "/":
        if num2 != 0 {
            fmt.Println("Result:", num1/num2)
        } else {
            fmt.Println("Division by zero!")
        }
    default:
        fmt.Println("Invalid operator. Use +, -, *, or /.")
    }
}
