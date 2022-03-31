package main

import (
	"fmt"

	"github.com/btcsuite/btcutil"
)

func main() {
	amount,_ := btcutil.NewAmount(0.02413015)

	(amount.Format(btcutil.AmountSatoshi))

	fmt.Println((amount.Format(btcutil.AmountSatoshi)))
}