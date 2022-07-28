package block

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/samthehai/simpleblockchain/signature"
)

// Blockchain is database ordered, back-linked list to store the blocks, in order to
// quickly get the latest block in the chain and efficently get the block by its hash
type Blockchain struct {
	address         string
	chain           []*Block
	transactionPool []*Transaction
}

// NewBlockChain creates a new blockchain with genesis block
func NewBlockChain(address string) *Blockchain {
	return &Blockchain{address: address, chain: []*Block{NewGenesisBlock()}}
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

func (bc *Blockchain) AddTransaction(sender string, recipient string, value float32,
	senderPublicKey *ecdsa.PublicKey, sign *signature.Signature) (*Transaction, error) {
	tx := NewTransaction(sender, recipient, value)

	if sender == MINING_SENDER {
		bc.transactionPool = append(bc.transactionPool, tx)
		return tx, nil
	}

	if err := bc.VerifyTransactionSignature(senderPublicKey, sign, tx); err != nil {
		return nil, fmt.Errorf("verify transaction: %w", err)
	}

	return tx, nil
}

func (bc *Blockchain) VerifyTransactionSignature(senderPublicKey *ecdsa.PublicKey, sign *signature.Signature,
	tx *Transaction) error {
	m, _ := json.Marshal(tx)
	h := sha256.Sum256([]byte(m))
	if ecdsa.Verify(senderPublicKey, h[:], sign.R, sign.S) {
		return nil
	}

	return fmt.Errorf("invalid transaction signature")
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

func (bc *Blockchain) CalculateTotalAmount(address string) float32 {
	var totalAmount float32
	for _, b := range bc.chain {
		for _, tx := range b.Transactions {
			if tx.RecipientAddress == address {
				totalAmount += tx.Value
			}

			if tx.SenderAddress == address {
				totalAmount -= tx.Value
			}
		}
	}
	return totalAmount
}

func (bc *Blockchain) Print() {
	for i, b := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i,
			strings.Repeat("=", 25))
		b.Print()
	}
	fmt.Printf("%s\n\n", strings.Repeat("*", 60))
}
