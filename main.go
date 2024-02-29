package main

import (
	"log"
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

	// w := wallet.NewWallet()

	// fmt.Println(w.PrivateKeyStr())
	// fmt.Println(w.PublicKeyStr())
	// fmt.Println(w.BlockchainAddress())

	// t := wallet.NewTransaction(w.PrivateKey(), w.PublicKey(), w.BlockchainAddress(), "B", 1.0)

	// fmt.Printf("signature %s", t.GenerateSignature())

	// walletM := wallet.NewWallet()
	// walletA := wallet.NewWallet()
	// walletB := wallet.NewWallet()

	// // wallet
	// t := wallet.NewTransaction(walletA.PrivateKey(), walletA.PublicKey(), walletA.BlockchainAddress(), walletB.BlockchainAddress(), 1.0)

	// // Blockchain

	// blockchain := block.NewBlockchain(walletM.BlockchainAddress())

	// isAdded := blockchain.AddTransaction(walletA.BlockchainAddress(), walletB.BlockchainAddress(), 1.0, walletA.PublicKey(), t.GenerateSignature())

	// fmt.Println("Added?", isAdded)

	// blockchain.Mining()
	// blockchain.Print()

	// fmt.Printf("A %.1f\n", blockchain.CalculateTotalAmount(walletA.BlockchainAddress()))
	// fmt.Printf("B %.1f\n", blockchain.CalculateTotalAmount(walletB.BlockchainAddress()))
	// fmt.Printf("M %.1f\n", blockchain.CalculateTotalAmount(walletM.BlockchainAddress()))

}

// 51 attack
