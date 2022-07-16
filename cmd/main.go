package main

import "github.com/samthehai/simpleblockchain"

func main() {
	blockchain := simpleblockchain.NewBlockChain()
	blockchain.AddBlock("send 1 BTC to No-One")
	blockchain.AddBlock("send 2 BTC to Some-One")
	blockchain.Print()
}
