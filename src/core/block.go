package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
	"time"
)

// Block 区块
type Block struct {
	Timestamp    int64           // 区块创建时间戳
	Transactions []*Transaction  // 交易数据
	PreBlockHash []byte          // 前一个区块的哈希值
	Hash         []byte          // 区块自身的哈希值，用于校验区块数据是否有效
	Nonce		 int             // 工作量证明使用的数值
}

func NewBlock(transactions []*Transaction, preBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), transactions, preBlockHash, []byte{}, 0}
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
func NewGenesisBlock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

/*
Serialize 序列化区块成字节数组
 */
func (block *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(block)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

/*
Deserialize 反序列化区块
*/
func Deserialize(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}

/*
HashTransactions 计算区块里所有交易的哈希值
 */
func (block *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range block.Transactions {
		txHashes = append(txHashes, tx.ID)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))

	return txHash[:]
}