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

	bc := NewBlockchain()

	bc.AddBlock("First Block")
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
