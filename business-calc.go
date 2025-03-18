package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"os"
	"runtime"
	"strconv"
)

// Reset the colour
var resetColour = "\033[0m"

// Text colours
var textRed = "\033[31m" 
var textGreen = "\033[32m" 
var textYellow = "\033[33m" 
var textBlue = "\033[34m" 
var textMagenta = "\033[35m" 
var textCyan = "\033[36m"
var textWhite = "\033[37m"
var textBoldWhite = "\033[1;37m"  
var textBlack = "\033[30m"
var textBoldBlack = "\033[1;30m"

// Background colours
var bgRed="\033[41m"
var bgGreen="\033[42m"
var bgYellow="\033[43m"
var bgBlue="\033[44m"
var bgPurple="\033[45m"
var bgCyan="\033[46m"
var bgWhite="\033[47m"
var bgBlack="\033[40m"  

func mainMenu() {
	fmt.Println("")
	fmt.Println("   "+textBoldBlack+"Press any button to continue..."+resetColour)
	fmt.Scanln()
	calc()
}

func divisorZero() {
	fmt.Println("")
	fmt.Println("   "+bgRed+textBoldWhite+" Any number divided by 0 is undefined in mathematics "+resetColour)
	mainMenu()
}

func wrongInput(input string) {
	fmt.Println("")
	fmt.Println("   "+bgRed+textBoldWhite+" "+input+" must be a number "+resetColour)
	mainMenu()
} 

func calc() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("")
	fmt.Println("   "+textBoldBlack+"Please select a number from the following options: "+resetColour)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("   "+bgYellow+textBoldWhite+" 1) To calculate Acid Test Ratio         "+resetColour)
	fmt.Println("")
	fmt.Println("   "+bgBlue+textBoldWhite+" 2) To calculate Current Assets          "+resetColour)
	fmt.Println("")
	fmt.Println("   "+bgPurple+textBoldWhite+" 3) To calculate Inventry Holding Period "+resetColour)
	fmt.Println("")
	fmt.Println("   "+bgCyan+textBoldWhite+" 4) To calculate Gross Profit Margin     "+resetColour)
	fmt.Println("")
	fmt.Println("   "+bgWhite+textBoldBlack+" 5) To calculate Net Profit Margin       "+resetColour)
	fmt.Println("")
	fmt.Println("   "+bgBlack+textBoldWhite+" 6) To calculate Current Ratio           "+resetColour)
	fmt.Println("")
	fmt.Println("   "+bgRed+textBoldWhite+" 7) Exit Program                         "+resetColour)
	fmt.Println("")
	fmt.Println("")
	fmt.Print("   "+textBoldBlack+"Please enter an option: ")

	var option int
	fmt.Scan(&option)
	fmt.Println(""+resetColour)

	var num1 string
	var num2 string
	var num3 string

	if option == 1 {
		// Acid Test Ratio
		fmt.Print("\033[H\033[2J")
		fmt.Println("")
		fmt.Println("   "+bgYellow+"                                                             "+resetColour)
		fmt.Println("   "+bgYellow+textBoldWhite+"                  Acid Test Ratio Calculater                 "+resetColour)
		fmt.Println("   "+bgYellow+textBoldWhite+"     [(current assets - inventry) / current liabilities]     "+resetColour)
		fmt.Println("   "+bgYellow+"                                                             "+resetColour)
		fmt.Println(""+textBlack)
		// Current Assets
		fmt.Print("   Please enter current assets: ")
		fmt.Scan(&num1)
		num1Float64, _ := strconv.ParseFloat(num1, 64)
		validateNum1 := validator.New()
		validateNum1Err := validateNum1.Var(num1, "numeric")
		if validateNum1Err != nil {
			input := "Current assets"
			wrongInput(input)                        		
		}
		// Inventry
		fmt.Print("   Please enter inventry: ")
		fmt.Scan(&num2)
		num2Float64, _ := strconv.ParseFloat(num2, 64)
		validateNum2 := validator.New()
		validateNum2Err := validateNum2.Var(num2, "numeric")
		if validateNum2Err != nil {
			input := "Inventry"
			wrongInput(input)
		}
		// Current Liabilities
		fmt.Print("   Please enter current liabilities: ")
		fmt.Scan(&num3)
		num3Float64, _ := strconv.ParseFloat(num3, 64)
		validateNum3 := validator.New()
		validateNum3Err := validateNum3.Var(num3, "numeric")
		if validateNum3Err != nil {
			input := "Current liabilities"
			wrongInput(input)			
		} else if num3Float64 == 0 {
			divisorZero()
		}
		fmt.Println(""+resetColour)
		ans := (num1Float64 - num2Float64) / num3Float64
		fmt.Println("   "+bgGreen+textBoldWhite+" The answer is:",ans,resetColour)
		mainMenu()
	} else if option == 2 {
		// Current Assets
		fmt.Print("\033[H\033[2J")
		fmt.Println("")
		fmt.Println("   "+bgBlue+"                                                             "+resetColour)
		fmt.Println("   "+textBoldWhite+bgBlue+"                  Current Assets Calculater                  "+resetColour)
		fmt.Println("   "+textBoldWhite+bgBlue+"     [(current assets + inventry) / current liabilities]     "+resetColour)
		fmt.Println("   "+bgBlue+"                                                             "+resetColour)
		fmt.Println(""+textBlack)
		// Current Assets
		fmt.Print("   Please enter current assets: ")
		fmt.Scan(&num1)
		num1Float64, _ := strconv.ParseFloat(num1, 64)
		validateNum1 := validator.New()
		validateNum1Err := validateNum1.Var(num1, "numeric")
		if validateNum1Err != nil {
			input := "Current assets"
			wrongInput(input)                        			
		}
		// Inventry
		fmt.Print("   Please enter inventry: ")
		fmt.Scan(&num2)
		num2Float64, _ := strconv.ParseFloat(num2, 64)
		validateNum2 := validator.New()
		validateNum2Err := validateNum2.Var(num2, "numeric")
		if validateNum2Err != nil {
			input := "Inventry"
			wrongInput(input)                        			
		}
		// Current Liabilities
		fmt.Print("   Please enter current liabilities: ")
		fmt.Scan(&num3)
		num3Float64, _ := strconv.ParseFloat(num3, 64)
		validateNum3 := validator.New()
		validateNum3Err := validateNum3.Var(num3, "numeric")
		if validateNum3Err != nil {
			input := "Current liabilities"
			wrongInput(input)                        			
		} else if num3Float64 == 0 {
			divisorZero()
		}		
		fmt.Println(""+resetColour)
		ans := (num1Float64 + num2Float64) / num3Float64
		fmt.Println("   "+bgGreen+textBoldWhite+" The answer is:",ans,resetColour)
		mainMenu()
	} else if option == 3 {
		//Inventry Holding Period
		fmt.Print("\033[H\033[2J")
		fmt.Println("")
		fmt.Println("   "+bgPurple+"                                                             "+resetColour)
		fmt.Println("   "+bgPurple+textBoldWhite+"              Inventry Holding Period Calculater             "+resetColour)
		fmt.Println("   "+bgPurple+textBoldWhite+"          [(closing inventry / cost of sales) X 365]         "+resetColour)
		fmt.Println("   "+bgPurple+"                                                             "+resetColour)
		fmt.Println(""+textBlack)
		// Closing Inventory
		fmt.Print("   Please enter closing inventory: ")
		fmt.Scan(&num1)
		num1Float64, _ := strconv.ParseFloat(num1, 64)
		validateNum1 := validator.New()
		validateNum1Err := validateNum1.Var(num1, "numeric")
		if validateNum1Err != nil {
			input := "Closing inventory"
			wrongInput(input)                        		
		}
		// Inventry
		fmt.Print("   Please enter cost of sales: ")
		fmt.Scan(&num2)
		num2Float64, _ := strconv.ParseFloat(num2, 64)
		validateNum2 := validator.New()
		validateNum2Err := validateNum2.Var(num2, "numeric")
		if validateNum2Err != nil {
			input := "Cost of sales"
			wrongInput(input)                        			
		} else if num2Float64 == 0 {
		        divisorZero()
		}
		fmt.Println(""+resetColour)
		ans := (num1Float64 / num2Float64) * 365
		fmt.Println("   "+bgGreen+textBoldWhite+" The answer is:",ans,resetColour)
		mainMenu()
	} else if option == 4 {
		//Gross Profit Margin
		fmt.Print("\033[H\033[2J")
		fmt.Println("")
		fmt.Println("   "+bgCyan+"                                                             "+resetColour)
		fmt.Println("   "+bgCyan+textBoldWhite+"                Gross Profit Margin Calculater               "+resetColour)
		fmt.Println("   "+bgCyan+textBoldWhite+"            [(gross profit / sales revenue) X 100]           "+resetColour)
		fmt.Println("   "+bgCyan+"                                                             "+resetColour)
		fmt.Println(""+textBlack)
		// Gross Profit
		fmt.Print("   Please enter gross profit: ")
		fmt.Scan(&num1)
		num1Float64, _ := strconv.ParseFloat(num1, 64)
		validateNum1 := validator.New()
		validateNum1Err := validateNum1.Var(num1, "numeric")
		if validateNum1Err != nil {
			input := "Gross profit"
			wrongInput(input)
		}
		// Inventry
		fmt.Print("   Please enter sales revenue: ")
		fmt.Scan(&num2)
		num2Float64, _ := strconv.ParseFloat(num2, 64)
		validateNum2 := validator.New()
		validateNum2Err := validateNum2.Var(num2, "numeric")
		if validateNum2Err != nil {
			input := "Sales revenue"
			wrongInput(input)                        			
		} else if num2Float64 == 0 {
		        divisorZero()                		
		}
		fmt.Println(""+resetColour)
		ans := (num1Float64 / num2Float64) * 100
		fmt.Println("   "+bgGreen+textBoldWhite+" The answer is:",ans,resetColour)
		mainMenu()
	} else if option == 5 {
		// Net Profit Margin
		fmt.Print("\033[H\033[2J")
		fmt.Println("")
		fmt.Println("   "+bgWhite+"                                                             "+resetColour)
		fmt.Println("   "+bgWhite+textBoldBlack+"                Gross Profit Margin Calculater               "+resetColour)
		fmt.Println("   "+bgWhite+textBoldBlack+"             [(net profit / sales revenue) X 100]            "+resetColour)
		fmt.Println("   "+bgWhite+"                                                             "+resetColour)
		fmt.Println(""+textBlack)
		// Net Profit
		fmt.Print("   Please enter net profit: ")
		fmt.Scan(&num1)
		num1Float64, _ := strconv.ParseFloat(num1, 64)
		validateNum1 := validator.New()
		validateNum1Err := validateNum1.Var(num1, "numeric")
		if validateNum1Err != nil {
			input := "Net profit"
			wrongInput(input)                        			
		}
		// Sales Revenue
		fmt.Print("   Please enter sales revenue: ")
		fmt.Scan(&num2)
		num2Float64, _ := strconv.ParseFloat(num2, 64)
		validateNum2 := validator.New()
		validateNum2Err := validateNum2.Var(num2, "numeric")
		if validateNum2Err != nil {
			input := "Sales revenue"
			wrongInput(input)                        		
		} else if num2Float64 == 0 {
		        divisorZero()                		
		}		
		fmt.Println(""+resetColour)
		ans := (num1Float64 / num2Float64) * 100
		fmt.Println("   "+bgGreen+textBoldWhite+" The answer is:",ans,resetColour)
		mainMenu()
	} else if option == 6 {
		//Current Ratio
		fmt.Print("\033[H\033[2J")
		fmt.Println("")
		fmt.Println("   "+bgBlack+"                                                             "+resetColour)
		fmt.Println("   "+bgBlack+textBoldWhite+"                   Current Ratio Calculater                  "+resetColour)
		fmt.Println("   "+bgBlack+textBoldWhite+"            [current assets / current liabilities]           "+resetColour)
		fmt.Println("   "+bgBlack+"                                                             "+resetColour)
		fmt.Println(""+textBlack)
		// Net Profit
		fmt.Print("   Please enter current assets: ")
		fmt.Scan(&num1)
		num1Float64, _ := strconv.ParseFloat(num1, 64)
		validateNum1 := validator.New()
		validateNum1Err := validateNum1.Var(num1, "numeric")
		if validateNum1Err != nil {
			input := "Current assets"
			wrongInput(input)                        		
		}
		// Sales Revenue
		fmt.Print("   Please enter current liabilities: ")
		fmt.Scan(&num2)
		num2Float64, _ := strconv.ParseFloat(num2, 64)
		validateNum2 := validator.New()
		validateNum2Err := validateNum2.Var(num2, "numeric")
		if validateNum2Err != nil {
			input := "Current liabilities"
			wrongInput(input)                        			
		} else if num2Float64 == 0 {
		        divisorZero()                		
		}
		fmt.Println(""+resetColour)
		ans := num1Float64 / num2Float64
		fmt.Println("   "+bgGreen+textBoldWhite+" The answer is:",ans,resetColour)
		mainMenu()
	} else if option == 7 {
		os.Exit(0)
	} else {
		fmt.Println("")
		fmt.Println("   "+bgRed+textBoldWhite+" The input must be a number between 1-7 "+resetColour)
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
