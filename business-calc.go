package main

import (
        "fmt"
        "github.com/go-playground/validator/v10"
        "runtime"
        "strconv"
)

func mainMenu() {
        fmt.Println("")
        fmt.Println("   Press any button to continue...")
        fmt.Scanln()
        calc()
}

func calc() {
        fmt.Print("\033[H\033[2J")
        fmt.Println("")
        fmt.Println("   Please select a number from the following options:")
        fmt.Println("")
        fmt.Println("")                                                                                                              
        fmt.Println("   1) To calculate Acid Test Ratio")
        fmt.Println("")
        fmt.Println("   2) To calculate Current Assets")
        fmt.Println("")
        fmt.Println("   3) To calculate Inventry Holding Period")
        fmt.Println("")
        fmt.Println("   4) To calculate Gross Profit Margin")
        fmt.Println("")
        fmt.Println("   5) To calculate Current Ratio")
        fmt.Println("")
        fmt.Println("   6) Exit Program")
        fmt.Println("")
        fmt.Println("")
        fmt.Print("   Please enter option: ")

        var option int
        fmt.Scan(&option)

        var num1 string
        var num2 string
        var num3 string

        if option == 1 {
                fmt.Print("\033[H\033[2J")
                fmt.Println("")
                fmt.Println("              Acid Test Ratio Calculater")
                fmt.Println("   (current assets - inventry / current liabilities)")
                fmt.Println("")
                // Current Assets
                fmt.Print("   Please enter current assets: ")
                fmt.Scan(&num1)
                num1Float64, _ := strconv.ParseFloat(num1, 64)
                validateNum1 := validator.New()
                validateNum1Err := validateNum1.Var(num1, "numeric")
                if validateNum1Err != nil {
                        fmt.Println("   Current assets must be a number")
                        mainMenu()
                }
                // Inventry
                fmt.Print("   Please enter inventry: ")
                fmt.Scan(&num2)
                num2Float64, _ := strconv.ParseFloat(num2, 64)
                validateNum2 := validator.New()
                validateNum2Err := validateNum2.Var(num2, "numeric")
                if validateNum2Err != nil {
                        fmt.Println("   Inventry must be a number")
                        mainMenu()
                }
                // Current Liabilities
                fmt.Print("   Please enter current liabilities: ")
                fmt.Scan(&num3)
                num3Float64, _ := strconv.ParseFloat(num3, 64)
                validateNum3 := validator.New()
                validateNum3Err := validateNum3.Var(num3, "numeric")
                if validateNum3Err != nil {
                        fmt.Println("   Current liabilities must be a number")
                        mainMenu()
                }
                fmt.Println("")
                ans := num1Float64 - num2Float64/num3Float64
                fmt.Println("   The answer is: ", ans)
                mainMenu()

        } else if option == 2 {

        } else if option == 3 {

                } else if option == 4 {

        } else if option == 5 {

                } else if option == 6 {

                } else {
                fmt.Println("")
                fmt.Println("   The input must be a number between 1-6")
                mainMenu()
        }
}

func main() {
        if runtime.GOOS != "linux" {
                fmt.Println("Operating system must be GNU/Linux to work")
        } else {
                calc()
        }
}
