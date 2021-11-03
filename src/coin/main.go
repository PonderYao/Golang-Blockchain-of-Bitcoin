package main

import (
	"core"
	"fmt"
	"strconv"
)

func main() {
	// 初始化区块链，创建第一个区块（创世纪区块）
	blockchain := core.NewBlockchain()

	blockchain.AddBlock("Send 1 BTC to Ivan")
	blockchain.AddBlock("Send 2 BTC to Ivan")

	for _, block := range blockchain.Blocks {
		fmt.Printf("Prev.Hash: %x\n", block.PreBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		// 工作量证明
		pow := core.NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
