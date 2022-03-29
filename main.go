package main

import (
	"fmt"
	"strconv"

	"github.com/CryptocurrencyCabal/btcutil"
	"github.com/blockcypher/gobcy"
)

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

	fresult, _ := strconv.ParseFloat(string(result), 64)
	
	amount := btcutil.Amount(fresult)
	fmt.Println("amount: ", amount)
}