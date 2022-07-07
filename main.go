package main

import (
	"fmt"
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
)

var infura = "https://mainnet.infura.io/v3/5aa06cf6c9d54a79bc1ce4f8ca6cfa6d"

func main() {

	client, err := ethclient.DialContext(context.Background(), infura)
	if err != nil {
		fmt.Println("error logged during connection:", err)
	}
	// in other to avoid a memory lick
	defer client.Close()
}
