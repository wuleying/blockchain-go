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

	bc.AddBlock("1")
	bc.AddBlock("2")

	for _, block := range bc.blocks {
		clog.Info("Prev. hash: %x", block.PrevBlockHash)
		clog.Info("Data: %s", block.Data)
		clog.Info("Hash: %x", block.Hash)
		pow := NewProofOfWork(block)
		clog.Info("PoW: %s", strconv.FormatBool(pow.Validate()))
	}
}
