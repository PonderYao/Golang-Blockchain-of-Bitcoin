package core

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Client struct {}

func (client *Client) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  createblockchain -address ADDRESS - Create a blockchain and send genesis block reward to ADDRESS")
	fmt.Println("  getbalance -address ADDRESS - Get balance of ADDRESS")
	fmt.Println("  send -from FROM -to TO -amount AMOUNT - Send AMOUNT of coins from FROM address to TO")
	fmt.Println("  printchain - Print all the blocks of the blockchain")
}

func (client *Client) validateArgs() {
	if len(os.Args) < 2 {
		client.printUsage()
		os.Exit(1)
	}
}

func (client *Client) printChain() {
	blockchain := NewBlockchain("")
	defer blockchain.db.Close()

	iterator := blockchain.Iterator()

	// 从后往前遍历
	for {
		block := iterator.Next()
		fmt.Printf("Prev.hash: %x\n", block.PreBlockHash)
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
createBlockchain 创建区块链
 */
func (client *Client) createBlockchain(address string) {
	blockchain := CreateBlockchain(address)
	blockchain.db.Close()
	fmt.Println("Done!")
}

/*
getBalance 获取余额
 */
func (client *Client) getBalance(address string) {
	blockchain := NewBlockchain(address)
	defer blockchain.db.Close()

	balance := 0
	UTXOs := blockchain.FindUTXO(address)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance of '%s': %d\n", address, balance)
}

/*
send 发送比特币
 */
func (client *Client) send(from, to string, amount int) {
	blockchain := NewBlockchain(from)
	defer blockchain.db.Close()

	tx := NewUTXOTransaction(from, to, amount, blockchain)
	blockchain.MineBlock([]*Transaction{tx})
	fmt.Println("Success!")
}

/*
Run 解析命令行的参数并执行
 */
func (client *Client) Run() {
	client.validateArgs()

	createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	createBlockchainAddress := createBlockchainCmd.String("address", "", "The address to create blockchain")
	getBalanceAddress := getBalanceCmd.String("address", "", "The address to get balance")
	sendFrom := sendCmd.String("from", "", "Source wallet address")
	sendTo := sendCmd.String("to", "", "Destination wallet address")
	sendAmount := sendCmd.Int("amount", 0, "Amount to send")

	switch os.Args[1] {
	case "createblockchain":
		err := createBlockchainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "getbalance":
		err := getBalanceCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "send":
		err := sendCmd.Parse(os.Args[2:])
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

	if createBlockchainCmd.Parsed() {
		if *createBlockchainAddress == "" {
			createBlockchainCmd.Usage()
			os.Exit(1)
		}
		client.createBlockchain(*createBlockchainAddress)
	}

	if getBalanceCmd.Parsed() {
		if *getBalanceAddress == "" {
			getBalanceCmd.Usage()
			os.Exit(1)
		}
		client.getBalance(*getBalanceAddress)
	}

	if sendCmd.Parsed() {
		if *sendFrom == "" || *sendTo == "" || *sendAmount <= 0 {
			sendCmd.Usage()
			os.Exit(1)
		}
		client.send(*sendFrom, *sendTo, *sendAmount)
	}

	if printChainCmd.Parsed() {
		client.printChain()
	}
}
