package main

import (
	"fmt"
	"log"

	"github.com/samthehai/simpleblockchain/app"
	"github.com/samthehai/simpleblockchain/block"
	"github.com/samthehai/simpleblockchain/server"
	blockchainserver "github.com/samthehai/simpleblockchain/server/blockchain"
	walletserver "github.com/samthehai/simpleblockchain/server/wallet"
	"github.com/samthehai/simpleblockchain/wallet"
)

func main() {
	createSeed()
	blockchainServer := blockchainserver.NewBlockChainServer(8080)
	walletServer := walletserver.NewWalletServer(8081, "http://127.0.0.1:8080")

	servers := []server.Server{blockchainServer, walletServer}
	app := app.NewApp(servers)
	log.Fatal(app.Run())
}

func createSeed() {
	walletM := wallet.NewWallet()
	walletA := wallet.NewWallet()
	walletB := wallet.NewWallet()

	tx := wallet.NewTransaction(walletA.PrivateKey, walletA.PublicKey,
		walletA.BlockchainAddress, walletB.BlockchainAddress, 1.0)

	blockchain := block.NewBlockChain(walletM.BlockchainAddress)
	blockchain.AddTransaction(walletA.BlockchainAddress, walletB.BlockchainAddress, 1.0,
		walletA.PublicKey, tx.GenerateSignature())

	blockchain.Mining()

	fmt.Printf("A %.1f\n", blockchain.CalculateTotalAmount(walletA.BlockchainAddress))
	fmt.Printf("B %.1f\n", blockchain.CalculateTotalAmount(walletB.BlockchainAddress))
	fmt.Printf("M %.1f\n", blockchain.CalculateTotalAmount(walletM.BlockchainAddress))
}
