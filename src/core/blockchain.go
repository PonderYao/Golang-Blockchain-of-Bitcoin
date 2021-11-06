package core

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"
const genesisKey = "leader"

// Blockchain 区块链
type Blockchain struct {
	tip []byte
	Db  *bolt.DB
}

type BlockchainIterator struct {
	currentHash []byte
	Db          *bolt.DB
}

func NewBlockchain() *Blockchain {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))

		if bucket == nil {
			fmt.Println("No existing blockchain found. Creating a new one")
			genesis := NewGenesisBlock()

			bucket, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				log.Panic(err)
			}

			err = bucket.Put(genesis.Hash, genesis.Serialize())
			if err != nil {
				log.Panic(err)
			}

			err = bucket.Put([]byte(genesisKey), genesis.Hash)
			if err != nil {
				log.Panic(err)
			}
			tip = genesis.Hash
		} else {
			tip = bucket.Get([]byte(genesisKey))
		}
		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	blockchain := Blockchain{tip, db}

	return &blockchain
}

/*
Iterator 迭代器，用于构造可遍历区块链的工具
 */
func (blockchain *Blockchain) Iterator() *BlockchainIterator {
	iterator := &BlockchainIterator{blockchain.tip, blockchain.Db}
	return iterator
}

/*
Next 迭代器的遍历方式
 */
func (iterator *BlockchainIterator) Next() *Block {
	var block *Block

	err := iterator.Db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		encodedBlock := bucket.Get(iterator.currentHash)
		block = Deserialize(encodedBlock)

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	iterator.currentHash = block.PreBlockHash

	return block
}

func (blockchain *Blockchain) AddBlock(data string) {
	var lastHash []byte

	err := blockchain.Db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		lastHash = bucket.Get([]byte(genesisKey))

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	newBlock := NewBlock(data, lastHash)

	err = blockchain.Db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		err := bucket.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}

		err = bucket.Put([]byte(genesisKey), newBlock.Hash)
		if err != nil {
			log.Panic(err)
		}

		blockchain.tip = newBlock.Hash

		return nil
	})

	if err != nil {
		log.Panic(err)
	}
}