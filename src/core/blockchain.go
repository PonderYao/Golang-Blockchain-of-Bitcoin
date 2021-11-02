package core

// Blockchain 区块链
type Blockchain struct {
	Blocks []*Block
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

/*
AddBlock 在区块链的末尾添加区块，需要提供上一个区块的哈希值，除了第一个区块（创世纪区块）
 */
func (blockchain *Blockchain) AddBlock(data string) {
	preBlock := blockchain.Blocks[len(blockchain.Blocks) - 1]
	newBlock := NewBlock(data, preBlock.Hash)
	blockchain.Blocks = append(blockchain.Blocks, newBlock)
}