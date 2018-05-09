package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

// ProofOfWork 工作量证明结构体
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

// NewProofOfWork 创建并返回一个工作量证明结构体
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-MineTargetBits))

	pow := &ProofOfWork{b, target}

	return pow
}

// prepareData 装填数据
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	// 数据格式：前一区块hash+当前区块数据+16进制时间戳+16进制挖矿难度系数+16进制nonce
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(MineTargetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

// Run 运行工作量证明
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)

	// 验证工作量证明 nonce 不能大于 int64 最大值
	for nonce < MaxNonce {
		data := pow.prepareData(nonce)

		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			// 不正确，nonce加1
			nonce++
		}
	}

	fmt.Println()

	return nonce, hash[:]
}

// Validate 验证工作量证明
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}
