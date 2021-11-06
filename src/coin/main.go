package main

import (
	"core"
)

func main() {
	blockchain := core.NewBlockchain()
	defer blockchain.Db.Close()

	client := core.Client{blockchain}
	client.Run()
}
