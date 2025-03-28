package main_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
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
	auth := bind.NewKeyedTransactor(key, chainID)

	sim := simulated.NewBackend(map[common.Address]types.Account{
		auth.From: {Balance: big.NewInt(9e18)},
	})

	ca, tx, err := bind.DeployContract(auth, common.FromHex(lib.CertMetaData.Bin), sim.Client(), nil)
	if err != nil {
		t.Fatalf("Failed to deploy smart contract: %v", err)
	}
	assert.NotNil(t, tx)

	sim.Commit()

	id := "101"
	ct := lib.CertificatesOutput{
		Name:   "Deren",
		Course: "MBCC",
		Grade:  "S",
		Date:   "2025-03-28",
	}

	cert := lib.NewCert()
	instance := cert.Instance(sim.Client(), ca)

	tx, err = bind.Transact(instance, auth, cert.PackIssue(id, ct.Name, ct.Course, ct.Grade, ct.Date))
	if err != nil {
		t.Fatalf("Failed to call Issue method: %v", err)
	}
	assert.NotNil(t, tx)

	sim.Commit()

	val, err := bind.Call(instance, nil, cert.PackCertificates(id), cert.UnpackCertificates)
	if err != nil {
		t.Fatalf("Failed to call Certificates method: %v", err)
	}

	if assert.NotNil(t, val) {
		assert.EqualValues(t, val, ct)
	}
}
