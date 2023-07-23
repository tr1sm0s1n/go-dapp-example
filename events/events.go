package main

import (
    "context"
    "fmt"
    "log"

    "github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
    client, err := ethclient.Dial("ws://127.0.0.1:8545")
    if err != nil {
        log.Fatal(err)
    }

    contractAddress := common.HexToAddress("0x68042D342f4722a74b0C54Cb37ef3A7E8A57905d")
    query := ethereum.FilterQuery{
        Addresses: []common.Address{contractAddress},
    }

    logs := make(chan types.Log)
    sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
    if err != nil {
        log.Fatal(err)
    }

	fmt.Println("Listening for events...")

    for {
        select {
        case err := <-sub.Err():
            log.Fatal(err)
        case vLog := <-logs:
            fmt.Println(vLog)
        }
    }
}