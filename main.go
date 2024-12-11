package main

import "fmt"
import "os"
import "strconv"
import "errors"

const accountBalanceFileName = "balanceTxtFile.txt"

func ReadFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFileName)
	if err != nil {
		return 1000, errors.New("Failed to read file")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("Failed to parse file")
	}
	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprintf("%f", balance)
	os.WriteFile(accountBalanceFileName, []byte(balanceText), 0644)
}

func main() {
	var accountBalance,err = ReadFromFile()
	if err != nil{
		fmt.Println("Error")
		fmt.Println(err)
	}
	fmt.Println("Welcome to Go Bank!")
	for {

		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")
		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("Enter the amount to deposit: ")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount")
				continue
			} else {
				accountBalance += depositAmount
				fmt.Println("Your new balance is: ", accountBalance)
				writeBalanceToFile(accountBalance)
			}
		case 3:
			var withdrawAmount float64
			fmt.Print("Enter the amount to withdraw: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient funds")
			} else {
				accountBalance -= withdrawAmount
				fmt.Println("Your new balance is: ", accountBalance)
				writeBalanceToFile(accountBalance)
			}
		case 4:
			fmt.Println("Thank you for using Go Bank!")
		default:
			fmt.Println("Invalid choice")

		}

	}

}
