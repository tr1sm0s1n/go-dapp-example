package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/tr1sm0s1n/go-dapp-example/config"
	"github.com/tr1sm0s1n/go-dapp-example/lib"
	"github.com/tr1sm0s1n/go-dapp-example/middlewares"
)

type Certificate struct {
	ID     string `json:"id" binding:"required"`
	Name   string `json:"name" binding:"required"`
	Course string `json:"course" binding:"required"`
	Grade  string `json:"grade" binding:"required"`
	Date   string `json:"date" binding:"required"`
}

var cert = lib.NewCert()

func main() {
	config.LoadEnv()

	contract := os.Getenv("CONTRACT_ADDRESS")
	printContract := fmt.Sprintf("Contract: %s", contract)
	fmt.Println(printContract)

	contractAddress := common.HexToAddress(contract)
	client := config.DialClient(false)

	instance := cert.Instance(client, contractAddress)

	router := gin.Default()
	router.POST("/issue", func(ctx *gin.Context) {
		issueCertificate(ctx, client, instance)
	})
	router.GET(("/fetch/:id"), func(ctx *gin.Context) {
		fetchCertificate(ctx, instance)
	})
	router.GET(("/events"), func(ctx *gin.Context) {
		getEvents(ctx, contractAddress, client)
	})
	router.Run("localhost:8080")
}

func issueCertificate(ctx *gin.Context, client *ethclient.Client, instance *bind.BoundContract) {
	var newCertificate Certificate

	if err := ctx.ShouldBind(&newCertificate); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	auth := middlewares.AuthGenerator(client)
	trx, err := bind.Transact(instance, auth,
		cert.PackIssue(newCertificate.ID,
			newCertificate.Name,
			newCertificate.Course,
			newCertificate.Grade,
			newCertificate.Date),
	)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, trx)
}

func fetchCertificate(ctx *gin.Context, instance *bind.BoundContract) {
	param := ctx.Param("id")

	result, err := bind.Call(instance, nil, cert.PackCertificates(param), cert.UnpackCertificates)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, result)
}

func getEvents(ctx *gin.Context, contractAddress common.Address, client *ethclient.Client) {
	latest, _ := client.BlockNumber(context.Background())
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		ToBlock:   big.NewInt(int64(latest)),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	logIssuedSig := []byte("Issued(string,string,string)")
	logIssuedSigHash := crypto.Keccak256Hash(logIssuedSig)
	logIssuedEvents := []types.Log{}

	for _, vLog := range logs {
		switch vLog.Topics[0].Hex() {
		case logIssuedSigHash.Hex():
			logIssuedEvents = append(logIssuedEvents, vLog)
		}
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"events": logIssuedEvents})
}
