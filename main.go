package main

import (
	"math/big"
	"math"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"os"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/ethereum/go-ethereum/ethclient"
)

var ctx = context.Background()

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {
	// load .env file
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
  }
  
func main() {
	getBlockNumber()
	getBalanceAtBlock()
}

func getBlockNumber() {
	dotenv := goDotEnvVariable("INFURAURL")
	client, err := ethclient.DialContext(ctx, dotenv)
	if err != nil {
		fmt.Println("error logged during connection:", err)
	}
	// in other to avoid a memory lick
	defer client.Close()

	blockNum, err := client.BlockByNumber(ctx, nil)
	if err != nil {
		fmt.Println("error attempting block num:", err)
	}
	fmt.Println("Current Block Number: ", blockNum.Number())
}

func getBalanceAtBlock() {
	dotenv := goDotEnvVariable("INFURAURL")
	client, err := ethclient.DialContext(ctx, dotenv)
	if err != nil {
		fmt.Println("error logged during connection:", err)
	}
	// in other to avoid a memory lick
	defer client.Close()

	addr := "0xab3B229eB4BcFF881275E7EA2F0FD24eeaC8C83a"
	address := common.HexToAddress(addr)

	balance, err := client.BalanceAt(ctx, address, nil)
	if err != nil {
		fmt.Println("error attempting to get balance num:", err)
	}
	FBalance := new(big.Float)
	FBalance.SetString(balance.String())

	value := new(big.Float).Quo(FBalance, big.NewFloat(math.Pow10(18)))
	fmt.Println("Current user balance is: ", value)
}
