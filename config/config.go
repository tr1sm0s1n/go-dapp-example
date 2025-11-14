package config

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func DialClient(ws bool) *ethclient.Client {
	rawurl := "http://127.0.0.1:8545"
	if ws {
		rawurl = "ws://127.0.0.1:8545"
	}

	client, err := ethclient.Dial(rawurl)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
