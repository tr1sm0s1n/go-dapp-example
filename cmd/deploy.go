package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tr1sm0s1n/go-dapp-example/config"
	"github.com/tr1sm0s1n/go-dapp-example/lib"
	"github.com/tr1sm0s1n/go-dapp-example/middlewares"
)

func main() {
	config.LoadEnv()

	client := config.DialClient(false)
	contractAddress, trx, err := deployContract(client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Transaction has been committed!!")
	fmt.Println("--------------------------------")
	fmt.Printf("Contract Address: \033[32m%s\033[0m\n", contractAddress.String())
	fmt.Println("-----------------")
	fmt.Printf("Transaction Hash: \033[32m%s\033[0m\n", trx.Hash())
	fmt.Println("-----------------")
}

func deployContract(client *ethclient.Client) (common.Address, *types.Transaction, error) {
	auth := middlewares.AuthGenerator(client)
	contract, transaction, _, err := lib.DeployCert(auth, client)
	return contract, transaction, err
}
