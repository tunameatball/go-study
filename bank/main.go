package main

import (
	"fmt"
	"log"

	"github.com/go-study/bank/accounts"
)

func main() {
	account := accounts.NewAccount("kkh")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(account.Balance())
}
