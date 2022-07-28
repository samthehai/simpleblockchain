package main

import (
	"fmt"

	"github.com/samthehai/simpleblockchain/block"
	"github.com/samthehai/simpleblockchain/wallet"
)

func main() {
	walletM := wallet.NewWallet()
	walletA := wallet.NewWallet()
	walletB := wallet.NewWallet()

	tx := wallet.NewTransaction(walletA.PrivateKey, walletA.PublicKey,
		walletA.BlockchainAddress, walletB.BlockchainAddress, 1.0)

	blockchain := block.NewBlockChain(walletM.BlockchainAddress)
	blockchain.AddTransaction(walletA.BlockchainAddress, walletB.BlockchainAddress, 1.0,
		walletA.PublicKey, tx.GenerateSignature())

	blockchain.Mining()
	blockchain.Print()

	fmt.Printf("A %.1f\n", blockchain.CalculateTotalAmount(walletA.BlockchainAddress))
	fmt.Printf("B %.1f\n", blockchain.CalculateTotalAmount(walletB.BlockchainAddress))
	fmt.Printf("M %.1f\n", blockchain.CalculateTotalAmount(walletM.BlockchainAddress))
}
