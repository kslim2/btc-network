package main

import (
	"fmt"

	"github.com/blockcypher/gobcy"
)

type bal struct {
	content bool
	balance int
}

func main() {

	bc := gobcy.API{}
	bc.Token = "21044813c9c04e8499d5726f8d4721e0"
	bc.Coin = "btc" //options: "btc","bcy","ltc","doge","eth"
	bc.Chain = "test3" //depending on coin: "main","test3","test"

	balance, errBalance:= bc.GetAddrBal("mhNVRVjgh9bUBuZQ5dpMAhRSF2nQUBQsws", nil)
	if errBalance != nil {fmt.Println("error: ", errBalance)}

	fmt.Println("name: ", balance.Wallet.Name)
	fmt.Println("balance: ", balance.Balance)
	fmt.Println("final balance: ", balance.FinalBalance)
	
	result, _ := balance.FinalBalance.MarshalJSON()
	fmt.Println("wanted balance: ", string(result))
}