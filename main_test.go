package main_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/assert"
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
	assert.NotNil(t, tx)

	sim.Commit()

	id := "101"
	ct := struct {
		Name   string
		Course string
		Grade  string
		Date   string
	}{
		Name:   "Deren",
		Course: "MBCC",
		Grade:  "5",
		Date:   "2024-05-11",
	}

	tx, err = cert.Issue(auth, id, ct.Name, ct.Course, ct.Grade, ct.Date)
	if err != nil {
		t.Fatalf("Failed to call Issue method: %v", err)
	}
	assert.NotNil(t, tx)

	sim.Commit()

	val, err := cert.Certificates(nil, id)
	if err != nil {
		t.Fatalf("Failed to call Certificates method: %v", err)
	}

	if assert.NotNil(t, val) {
		assert.EqualValues(t, val, ct)
	}
}
