package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// Connect to the Ethereum network using an Ethereum client
	client, err := ethclient.Dial("https://goerli.infura.io/v3/")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum network: %v", err)
	}

	// Define the address to listen for transactions
	address := common.HexToAddress("0x2d87caaCAEa9C24FE1bfD4Fb22641077d4076f76")

	// Periodically poll the Ethereum network for new blocks
	for {
		// Get the latest block number
		blockNumber, err := client.BlockByNumber(context.Background(), nil)
		if err != nil {
			log.Fatalf("Failed to get latest block number: %v", err)
		}

		// Check each new block for transactions involving the specified address
		for i := blockNumber.Number().Uint64() - 1; i > 0; i-- {
			block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(i)))
			if err != nil {
				log.Fatalf("Failed to get block: %v", err)
			}

			for _, tx := range block.Transactions() {
				if tx.To() == nil {
					continue
				}
				if *tx.To() != address {
					continue
				}

				// Write a message to the console when a transaction is detected
				fmt.Printf("Transaction detected for address %s in block %d\n", address.String(), block.NumberU64())
			}
		}

		// Wait before polling the Ethereum network again
		time.Sleep(time.Second * 10)
	}
}
