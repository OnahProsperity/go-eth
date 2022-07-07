package main

import (
	"log"
	"os"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/ethereum/go-ethereum/ethclient"
)


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
}

func getBlockNumber() {
	dotenv := goDotEnvVariable("INFURAURL")
	client, err := ethclient.DialContext(context.Background(), dotenv)
	if err != nil {
		fmt.Println("error logged during connection:", err)
	}
	// in other to avoid a memory lick
	defer client.Close()

	blockNum, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		fmt.Println("error attempting block num:", err)
	}
	fmt.Println("Current Block Number: ", blockNum.Number())
}
