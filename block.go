package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Block 区块结构体
type Block struct {
	Timestamp     int64  // 当前时间戳，也就是区块创建的时间
	Data          []byte // 区块存储的实际有效的信息
	PrevBlockHash []byte // 存储的是前一个块的哈希
	Hash          []byte // 当前块的哈希
}

// SetHash
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

// NewBlock
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

// NewGenesisBlock 创世区块
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
