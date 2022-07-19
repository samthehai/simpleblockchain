package simpleblockchain

import (
	"fmt"
	"strings"
)

// Blockchain is database ordered, back-linked list to store the blocks, in order to
// quickly get the latest block in the chain and efficently get the block by its hash
type Blockchain struct {
	chain           []*Block
	transactionPool []*Transaction
}

// NewBlockChain creates a new blockchain with genesis block
func NewBlockChain() *Blockchain {
	return &Blockchain{chain: []*Block{NewGenesisBlock()}}
}

// AddBlock creates and adds new block to blockchain
func (bc *Blockchain) AddBlock(nonce uint64, prevHash [32]byte) *Block {
	prevBlock := bc.LastBlock()
	// Get transaction from transactionPool
	newBlock := NewBlock(nonce, prevBlock.Hash(), bc.transactionPool)
	// Add new block to chain
	bc.chain = append(bc.chain, newBlock)
	// Reset transactionPool
	bc.transactionPool = []*Transaction{}
	return newBlock
}

func (bc *Blockchain) AddTransaction(sender string, recipient string, value float32) *Transaction {
	tx := NewTransaction(sender, recipient, value)
	bc.transactionPool = append(bc.transactionPool, tx)
	return tx
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

func (bc *Blockchain) Print() {
	for i, b := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i,
			strings.Repeat("=", 25))
		b.Print()
	}
	fmt.Printf("%s\n\n", strings.Repeat("*", 60))
}
