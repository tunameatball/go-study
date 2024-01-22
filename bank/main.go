package main

import (
	"fmt"

	"github.com/go-study/bank/accounts"
)

func main() {
	account := accounts.NewAccount("kkh")
	account.Deposit(10)

	err := account.Withdraw(20)

	if err != nil {
		// log.Fatalln(err)
		fmt.Println(err)
	}

	fmt.Println(account.Balance(), account.Owner())
	fmt.Println(account.String())
}
