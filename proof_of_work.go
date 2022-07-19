package simpleblockchain

import (
	"fmt"
	"strings"
)

const MINING_DIFFICULTY = 3

func (bc *Blockchain) CopyTransactionPool() []*Transaction {
	transactions := make([]*Transaction, 0, len(bc.transactionPool))
	for _, t := range bc.transactionPool {
		transactions = append(transactions, NewTransaction(t.SenderAddress, t.RecipientAddress, t.Value))
	}
	return transactions
}

func (bc *Blockchain) validateProof(nonce uint64, previousHash [32]byte, transactions []*Transaction, difficulty int) bool {
	zeros := strings.Repeat("0", difficulty)
	guessBlock := Block{Header: BlockHeader{previousHash, 0, nonce}, Transactions: transactions}
	guessHashStr := fmt.Sprintf("%x", guessBlock.Hash())
	return guessHashStr[:difficulty] == zeros
}

func (bc *Blockchain) ProofOfWork() uint64 {
	transactions := bc.CopyTransactionPool()
	previousHash := bc.LastBlock().Hash()
	var nonce uint64 = 0
	for !bc.validateProof(nonce, previousHash, transactions, MINING_DIFFICULTY) {
		nonce += 1
	}
	return nonce
}
