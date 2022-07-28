package block

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

// BlockHeader is header of block
type BlockHeader struct {
	PrevBlockHash [32]byte `json:"prev_block_hash"` // hash of previous block
	Timestamp     int64    `json:"timestamp"`       // unix timestamp when the block is created
	Nonce         uint64   `json:"nonce"`           // value identified to solve the hash solution
}

// Block is the unit contains information (or transaction) of blockchain
type Block struct {
	Header       BlockHeader    `json:"header"`       // header of block
	Transactions []*Transaction `json:"transactions"` // the actual information is contained inside the block
}

// Hash uses to compute and assign Hash value for the block.
// The way hashes are calculated is very important feature in blockchain so that computational of
// hash should be difficult enough so that blockchain become secure preventing block be modified after
// be added.
// Use sha256 checksum algorithm to calculate the hash
func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256(m)
}

func (b *Block) Print() {
	fmt.Printf("timestamp         %d\n", b.Header.Timestamp)
	fmt.Printf("prev_block_hash   %x\n", b.Header.PrevBlockHash)
	fmt.Printf("hash              %x\n", b.Hash())
	for _, t := range b.Transactions {
		t.Print()
	}
}

// NewBlock creates and returns a new block base on data and previous block chain
func NewBlock(nonce uint64, prevBlockHash [32]byte, transactions []*Transaction) *Block {
	return &Block{
		Header:       BlockHeader{prevBlockHash, time.Now().Unix(), nonce},
		Transactions: transactions,
	}
}

// NewGenesisBlock creates and returns genesis block - the first block in blockchain
func NewGenesisBlock() *Block {
	return NewBlock(0, (&Block{}).Hash(), []*Transaction{})
}
