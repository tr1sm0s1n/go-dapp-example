package main_test

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/ethereum/go-ethereum/params"
	"github.com/tr1sm0s1n/go-dapp-example/lib"
)

func TestContract(t *testing.T) {
	key, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("Failed to generate key: %v", err)
	}

	chainID := params.AllDevChainProtocolChanges.ChainID

	auth, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		t.Fatalf("Failed to make transactor: %v", err)
	}

	sim := simulated.NewBackend(map[common.Address]types.Account{
		auth.From: {Balance: big.NewInt(9e18)},
	})

	_, tx, cert, err := lib.DeployCert(auth, sim.Client())
	if err != nil {
		t.Fatalf("Failed to deploy smart contract: %v", err)
	}

	fmt.Printf("Deploy pending: 0x%x\n", tx.Hash())

	sim.Commit()

	tx, err = cert.Issue(auth, "101", "Deren", "MBCC", "S", "11-05-2024")
	if err != nil {
		t.Fatalf("Failed to call Issue method: %v", err)
	}
	fmt.Printf("State update pending: 0x%x\n", tx.Hash())

	sim.Commit()

	val, err := cert.Certificates(nil, "101")
	if err != nil {
		t.Fatalf("Failed to call Certificates method: %v", err)
	}

	fmt.Printf("Value: %v\n", val)
}