package main

import (
	"fmt"
	"log"

	"github.com/blockchain-go/wallet"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	// myBlockchainAddress := "my_blockchain_address"
	// blockChain := NewBlockchain(myBlockchainAddress)
	// blockChain.Print()

	// blockChain.AddTransaction("A", "B", 1.0)
	// blockChain.Mining()
	// blockChain.Print()

	// blockChain.AddTransaction("C", "D", 2.0)
	// blockChain.AddTransaction("X", "Y", 3.0)
	// blockChain.Mining()
	// blockChain.Print()

	// fmt.Printf("C %1f\n", blockChain.CalculateTotalAmount("my_blockchain_address"))
	// fmt.Printf("C %1f\n", blockChain.CalculateTotalAmount("C"))
	// fmt.Printf("C %1f\n", blockChain.CalculateTotalAmount("D"))

	w := wallet.NewWallet()

	fmt.Println(w.PrivateKeyStr())
	fmt.Println(w.PublicKeyStr())
}
