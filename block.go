package main

import "time"

// Block 区块结构体
type Block struct {
	Id            int64
	Timestamp     int64
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

// NewBlock 创建新区块
func NewBlock(prevId int64, prevBlockHash []byte) *Block {
	id := prevId + 1
	block := &Block{id, time.Now().Unix(), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}
