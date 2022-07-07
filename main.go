package main

import (
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"math"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"log"
	"os"
	"io/ioutil"
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
	// getBlockNumber()
	// getBalanceAtBlock()
	// generateETHAddress()
	genKeyStoreAddress()
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
	infura := goDotEnvVariable("INFURAURL")
	client, err := ethclient.DialContext(ctx, infura)
	if err != nil {
		fmt.Println("error logged during connection:", err)
	}
	// in other to avoid a memory lick
	defer client.Close()

	addr := goDotEnvVariable("ADDRESS")
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

func generateETHAddress() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		fmt.Println("error logged during connection:", err)
	}
	publicKey := crypto.FromECDSA(privateKey)
	fmt.Println("Here is your new private key: ", privateKey, "and your public key: ", hexutil.Encode(publicKey))
	address := crypto.PubkeyToAddress(privateKey.PublicKey).Hex()

	fmt.Println("address: ", address)
}

func genKeyStoreAddress() {
	// key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "pasword"
	// theaddress, err := key.NewAccount(password)
	// if err != nil {
	// 	fmt.Println("error logged during connection:", err)
	// }
	// fmt.Println("address is: ", theaddress.Address)

	b, err := ioutil.ReadFile("./wallet/UTC--2022-07-07T20-10-12.830532744Z--7540d57e5f17fa967183c405d3d56c5e6ab1502a")
	if err != nil {
		log.Print(err)
	}
	key, err := keystore.DecryptKey(b, password)
	if err != nil {
		log.Print(err)
	}

	publicData := crypto.FromECDSA(key.PrivateKey)
	fmt.Println("the returned private data: ", hexutil.Encode(publicData))
	publicData = crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
	fmt.Println("the returned public key", hexutil.Encode(publicData))
	address := crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex()

	fmt.Println("address: ", address)
}
