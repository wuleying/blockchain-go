package main

import (
	"fmt"
	"github.com/go-clog/clog"
	"os"
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

	bc := NewBlockchain()

	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
