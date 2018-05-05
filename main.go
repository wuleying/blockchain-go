package main

import (
	"fmt"
	"github.com/go-clog/clog"
	"os"
	"strconv"
)

func init() {
	if err := clog.New(clog.CONSOLE, clog.ConsoleConfig{
		Level:      clog.INFO,
		BufferSize: 100},
	); err != nil {
		fmt.Printf("Init console log failed. error %+v.", err)
		os.Exit(1)
	}
}

func main() {
	defer clog.Shutdown()

	// 创建一个区块链
	bc := CreateBlockchain()

	// 第一个区块
	bc.AddBlock("First Block")
	// 第二个区块
	bc.AddBlock("Second Block")

	for _, block := range bc.blocks {
		clog.Info("Prev hash: %x", block.PrevBlockHash)
		clog.Info("Data: %s", block.Data)
		clog.Info("Hash: %x", block.Hash)
		clog.Info("Nonce: %d", block.Nonce)
		pow := NewProofOfWork(block)
		clog.Info("PoW: %s\n", strconv.FormatBool(pow.Validate()))
	}
}
