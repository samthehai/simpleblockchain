package simpleblockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

// Block is the unit contains information (or transaction) of blockchain
type Block struct {
	// Timestamp is the Unix timestamp when the block is created
	Timestamp int64
	// Data is the actual information is contained inside the block
	Data []byte
	// PrevBlockHash is the hash of previous block
	PrevBlockHash []byte
	// Hash is hash of current block
	Hash []byte
}

// SetHash uses to compute and assign Hash value for the block.
// The way hashes are calculated is very important feature in blockchain so that computational of
// hash should be difficult enough so that blockchain become secure preventing block be modified after
// be added.
// Use sha256 checksum algorithm to calculate the hash
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{timestamp, b.Data, b.PrevBlockHash}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

func (b *Block) Print() {
	fmt.Printf("timestamp         %d\n", b.Timestamp)
	fmt.Printf("data              %s\n", b.Data)
	fmt.Printf("prev_block_hash   %s\n", b.PrevBlockHash)
	fmt.Printf("hash              %s\n", b.Hash)
}

// NewBlock creates and returns a new block base on data and previous block chain
func NewBlock(data string, prevBlockHash []byte) *Block {
	newBlock := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	newBlock.SetHash()
	return newBlock
}

// NewGenesisBlock creates and returns genesis block - the first block in blockchain
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
