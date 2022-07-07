package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

var infuraURL = "wss://mainnet.infura.io/ws/v3/5aa06cf6c9d54a79bc1ce4f8ca6cfa6d"

func main() {
	// cxt := context.Background()
	client, err := ethclient.DialContext(context.Background(), infuraURL)
	if err != nil {
		fmt.Println("error logged during connection:", err)
	}
	// in other to avoid a memory lick
	defer client.Close()

	blockNum, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		fmt.Println("error attempting block num:", err)
	}
	fmt.Println(blockNum.Number())
}
