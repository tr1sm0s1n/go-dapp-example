package config

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
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

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
