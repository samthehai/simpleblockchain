package simpleblockchain

// Blockchain is database ordered, back-linked list to store the blocks, in order to
// quickly get the latest block in the chain and efficently get the block by its hash
type BlockChain struct {
	blocks []*Block
}

// NewBlockChain creates a new blockchain with genesis block
func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}

// AddBlock creates and adds new block to blockchain
func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func (bc *BlockChain) Print() {
	for _, b := range bc.blocks {
		b.Print()
	}
}
