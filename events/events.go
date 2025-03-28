package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/tr1sm0s1n/go-dapp-example/config"
	"github.com/tr1sm0s1n/go-dapp-example/lib"
)

func main() {
	config.LoadEnv()

	client := config.DialClient(true)
	contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Listening for events...")
	fmt.Println("-----------------------")

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case log := <-logs:
			event, _ := lib.NewCert().UnpackIssuedEvent(&log)
			rw, _ := json.Marshal(log)
			fmt.Println("Certificate issued!!")
			fmt.Println("--------------------")
			fmt.Printf("Course: \033[34m%s\033[0m\n", event.Course)
			fmt.Printf("ID: \033[34m%s\033[0m\n", event.Id)
			fmt.Printf("Grade: \033[34m%s\033[0m\n", event.Grade)
			fmt.Printf("Raw log: \033[32m%s\033[0m\n", string(rw))
			fmt.Println("--------------------")

		}
	}
}
