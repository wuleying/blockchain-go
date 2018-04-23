package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
	"time"
)

// ProofOfWork 工作量证明结构体
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

// NewProofOfWork
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-MINE_TARGET_BITS))
	pow := &ProofOfWork{b, target}
	return pow
}

// prepareData
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			IntToHex(int64(pow.block.Timestamp)),
			IntToHex(int64(MINE_TARGET_BITS)),
			IntToHex(int64(nonce)),
		}, []byte{})

	return data
}

// Run
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0
	startTime := time.Now()

	fmt.Printf("Mining the block containing [#%d]\n", pow.block.Id)

	for nonce < MAX_NONCE {
		data := pow.prepareData(nonce)

		hash = sha256.Sum256(data)
		fmt.Printf("\rhash: %x, nonce: %d", hash, nonce)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			fmt.Println()
			fmt.Printf("Mining time: [#%d] %s\n", pow.block.Id, time.Since(startTime))
			break
		} else {
			nonce++
		}
	}

	return nonce, hash[:]
}
