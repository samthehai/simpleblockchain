package main

import "github.com/samthehai/simpleblockchain"

func main() {
	blockchain := simpleblockchain.NewBlockChain("my_address")
	blockchain.Print()

	blockchain.AddTransaction("A", "B", 1.0)
	blockchain.Mining()
	blockchain.Print()

	blockchain.AddTransaction("C", "D", 2.0)
	blockchain.AddTransaction("X", "Y", 3.0)
	blockchain.Mining()
	blockchain.Print()

	blockchain.Print()
}
