package core

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Client struct {
	Bc *Blockchain
}

func (client *Client) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("	addblock -data BLOCK_DATA - add a block to the blockchain")
	fmt.Println("	printchain - print add the blocks of the blockchain")
}

func (client *Client) validateArgs() {
	if len(os.Args) < 2 {
		client.printUsage()
		os.Exit(1)
	}
}

func (client *Client) addBlock(data string) {
	client.Bc.AddBlock(data)
	fmt.Println("Success!")
}

func (client *Client) printChain() {
	iterator := client.Bc.Iterator()

	// 从后往前遍历
	for {
		block := iterator.Next()
		fmt.Printf("Prev.hash: %x\n", block.PreBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PreBlockHash) == 0 {
			break
		}
	}
}

/*
Run 解析命令行的参数并执行
 */
func (client *Client) Run() {
	client.validateArgs()

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	addBlockData := addBlockCmd.String("data", "", "Block data")

	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		client.printUsage()
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}
		client.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		client.printChain()
	}
}
