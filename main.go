package main

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"

	"github.com/CryptocurrencyCabal/btcutil"
	"github.com/blockcypher/gobcy"
)

func main() {
	
	bc := gobcy.API{}
	bc.Token = "21044813c9c04e8499d5726f8d4721e0"
	bc.Coin = "btc" //options: "btc","bcy","ltc","doge","eth"
	bc.Chain = "test3" //depending on coin: "main","test3","test"

	// getting the balance of an wallet
	balance, errBalance:= bc.GetAddrBal("mhNVRVjgh9bUBuZQ5dpMAhRSF2nQUBQsws", nil)
	if errBalance != nil {fmt.Println("error: ", errBalance)}

	fmt.Println("name: ", balance.Wallet.Name)
	fmt.Println("balance: ", balance.Balance)
	fmt.Println("final balance: ", balance.FinalBalance)

	// converting from satoshis to btc
	result, _ := balance.FinalBalance.MarshalJSON()

	fresult, _ := strconv.ParseFloat(string(result), 64)
	
	amount := btcutil.Amount(fresult)
	fmt.Println("amount: ", amount)
	fmt.Println()

	// check usage (the counter)
	usage, errUsage := bc.CheckUsage()
	if errUsage != nil {fmt.Println("error usage: ", errUsage)}

	fmt.Println("usage: ", usage.Limits.PerDay)
	fmt.Println("previous hits: ", len(usage.HitsHistory))
	fmt.Println()

	// get the latest unconfirmed transaction
	unTxs, errUnTxs := bc.GetUnTX()
	if errUnTxs != nil {fmt.Println("unconfirmed transaction error: ", errUnTxs)}
	fmt.Println("all unconfirmed transaction in form of list: ", unTxs)
	fmt.Println()

	// get a transaction information
	var transactionHash = "b888e06ff88e2cab8529c61c8ca5a60b2be9f404ad41a350ec569191fda8cfb9"
	tx, errTx := bc.GetTX(transactionHash, nil)
	if errTx != nil {fmt.Println("error tx: ", errTx)}
	fmt.Println("a transaction: ", tx)
	fmt.Println()

	// validator to define whether a transaction was confirmed
	// [0, 1] 1 for confirmed transaction, otherwise 0
	txConfig, errTxConfig := bc.GetTXConf(transactionHash)
	if errTxConfig != nil {fmt.Println("tx configuration: ", errTxConfig)}
	fmt.Println("txConfig: ", txConfig)
	fmt.Println()

	// create a new btc transaction 
	var inAddr = "mhNVRVjgh9bUBuZQ5dpMAhRSF2nQUBQsws"
	var outAddr = "mqDvHFjebVyEwRkWK1732CByQJHAgeMswf"
	var btcAmount = *big.NewInt(10000)
	txSkel := gobcy.TempNewTX(inAddr, outAddr, btcAmount)
	skel, errSkel := bc.NewTX(txSkel,  true)
	if errSkel != nil {fmt.Println("error creating a btc transaction: ", errSkel)}


	// sign a new btc transaction	
	var wif = "cTjBTn1u6mPtKErYWUZwD2Wabr4BuUTdtJP9cGLFM8N6q38a5u7w"
	w,_:= btcutil.DecodeWIF(wif)
	var privateKey = []string{hex.EncodeToString(w.PrivKey.Serialize())}
	errSignature := skel.Sign(privateKey)
	if errSignature != nil {fmt.Println("error signing transaction: ", errSignature)}
	fmt.Println()

	// sending a transaction
	sk, errSk := bc.SendTX(skel)
	if errSk != nil {fmt.Println("error broadcast a transaction: ", errSk)}
	fmt.Println("broadcasted transaction: ", sk)
	fmt.Println()

	fmt.Println("transaction hash: ", sk.Trans.Hash)
	fmt.Println()

	// get the chain 
	chain, errChain := bc.GetChain()
	if errChain != nil {fmt.Println("chain error: ", errChain)}

	fmt.Println("chain: ", chain)
	fmt.Println("chain height: ", chain.Height)
	fmt.Println()

	block, errBlock := bc.GetBlock(chain.Height, sk.Trans.BlockHash, nil)
	if errBlock != nil {fmt.Println("error block: ", errBlock)}

	fmt.Println("block: ", block)
	fmt.Println()
}