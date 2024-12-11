package main

import "fmt"

func main() {

	var accountBalance = 1000
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

		switch choice{
			case 1:
				fmt.Println("Your balance is: ", accountBalance)
			case 2:
				var depositAmount int
			  fmt.Print("Enter the amount to deposit: ")
			  fmt.Scan(&depositAmount)
		  	if depositAmount <= 0{
				   fmt.Println("Invalid amount")
				   continue
			  }else{
			  	accountBalance += depositAmount
			  	fmt.Println("Your new balance is: ", accountBalance)
			  }
			case 3:
			var withdrawAmount int
			fmt.Print("Enter the amount to withdraw: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount > accountBalance{
				fmt.Println("Insufficient funds")
			}else{
				accountBalance -= withdrawAmount
				fmt.Println("Your new balance is: ", accountBalance)
			}
			case 4:
			fmt.Println("Thank you for using Go Bank!")
			default:
			fmt.Println("Invalid choice")
			
		}
		
	}
	
	
}