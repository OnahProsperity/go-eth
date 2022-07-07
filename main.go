package main

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
)

var infura = "https://mainnet.infura.io/v3/5aa06cf6c9d54a79bc1ce4f8ca6cfa6d"

func main() {

	ethclient.DialContext(context.Background(), infura)
	// fmt.Println(a)

}
