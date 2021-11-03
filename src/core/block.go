package core

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Block 区块
type Block struct {
	Timestamp    int64   //区块创建时间戳
	Data         []byte  //区块包含的数据
	PreBlockHash []byte  //前一个区块的哈希值
	Hash         []byte  //区块自身的哈希值，用于校验区块数据是否有效
	Nonce		 int     //工作量证明使用的数值
}

func NewBlock(data string, preBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), preBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Nonce = nonce
	//block.SetHash()
	block.Hash = hash[:]
	return block
}

/*
NewGenesisBlock 创世纪区块
 */
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

/*
SetHash 计算并设置当前区块的哈希值
 */
func (block *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(block.Timestamp, 10))
	headers := bytes.Join([][]byte{block.PreBlockHash, block.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	block.Hash = hash[:]
}