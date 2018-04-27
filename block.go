package main

import (
	"time"
)

// Block 区块结构体
type Block struct {
	Timestamp     int64  // 当前区块创建时间
	Data          []byte // 区块存储的信息
	PrevBlockHash []byte // 前一个块的哈希
	Hash          []byte // 当前块的哈希
	Nonce         int    // Number once，在密码学中Nonce是一个只被使用一次的任意或非重复的随机数值
}

// NewBlock 创建并返回一个区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// NewGenesisBlock 创建并返回一个创世区块
func NewGenesisBlock() *Block {
	// 创世区块 前一个块的哈希为空
	return NewBlock("Genesis Block", []byte{})
}
