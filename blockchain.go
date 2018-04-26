package main

// Blockchain 区块链结构体
type Blockchain struct {
	blocks []*Block
}

// AddBlock 向区块链中添加一个新区块
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

// NewBlockchain 创建一个新区块链并生成一个创世区块
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
