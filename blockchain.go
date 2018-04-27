package main

// Blockchain 区块链结构体
type Blockchain struct {
	blocks []*Block
}

// AddBlock 向区块链中添加一个新区块
func (bc *Blockchain) AddBlock(data string) {
	// 获取区块链中最后一个块
	prevBlock := bc.blocks[len(bc.blocks)-1]
	// 添加新块
	newBlock := NewBlock(data, prevBlock.Hash)
	// 将新块加入区块链
	bc.blocks = append(bc.blocks, newBlock)
}

// NewBlockchain 创建一个新区块链并生成一个创世区块
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
