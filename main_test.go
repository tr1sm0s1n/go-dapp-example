package main

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/assert"
	"github.com/tr1sm0s1n/go-dapp-example/lib"
)

func simulateBackend() (*simulated.Backend, []*bind.TransactOpts, error) {
	k1, err := crypto.GenerateKey()
	if err != nil {
		return nil, nil, err
	}
	k2, err := crypto.GenerateKey()
	if err != nil {
		return nil, nil, err
	}

	chainID := params.AllDevChainProtocolChanges.ChainID
	t1 := bind.NewKeyedTransactor(k1, chainID)
	t2 := bind.NewKeyedTransactor(k2, chainID)

	sim := simulated.NewBackend(map[common.Address]types.Account{
		t1.From: {Balance: big.NewInt(9e18)},
		t2.From: {Balance: big.NewInt(9e18)},
	})

	return sim, []*bind.TransactOpts{t1, t2}, nil
}

func deployContract(sim *simulated.Backend, auth *bind.TransactOpts) (*common.Address, error) {
	contractAddress, _, err := bind.DeployContract(auth, common.FromHex(lib.CertMetaData.Bin), sim.Client(), nil)
	if err != nil {
		return nil, err
	}

	sim.Commit()
	return &contractAddress, nil
}

func TestDeploy(t *testing.T) {
	sim, auth, err := simulateBackend()
	if err != nil {
		t.Fatalf("Failed to create simulated backend: %v", err)
	}

	deployedContract, err := deployContract(sim, auth[0])
	if err != nil {
		t.Fatalf("Failed to deploy smart contract: %v", err)
	}

	contractAddress := crypto.CreateAddress(auth[0].From, 0)
	assert.Equal(t, contractAddress, *deployedContract)
}

func TestIssue(t *testing.T) {
	sim, auth, err := simulateBackend()
	if err != nil {
		t.Fatalf("Failed to create simulated backend: %v", err)
	}

	deployedContract, err := deployContract(sim, auth[0])
	if err != nil {
		t.Fatalf("Failed to deploy smart contract: %v", err)
	}

	cert := lib.NewCert()
	instance := cert.Instance(sim.Client(), *deployedContract)

	id := big.NewInt(21)
	ct := lib.CertificatesOutput{
		Name:   "Shin",
		Course: "Catalyst",
		Grade:  "S",
		Date:   "11-08-2025",
	}

	_, err = bind.Transact(instance, auth[0], cert.PackIssue(id, ct.Name, ct.Course, ct.Grade, ct.Date))
	if err != nil {
		t.Fatalf("Failed to issue certificate: %v", err)
	}

	sim.Commit()

	logs, err := sim.Client().FilterLogs(context.Background(), ethereum.FilterQuery{
		Addresses: []common.Address{*deployedContract},
	})
	if err != nil {
		t.Fatalf("Failed to filter logs: %v", err)
	}
	assert.NotEmpty(t, logs)

	event, err := cert.UnpackIssuedEvent(&logs[0])
	if err != nil {
		t.Fatalf("Failed to unpack 'Issued' event: %v", err)
	}

	assert.Equal(t, crypto.Keccak256Hash([]byte(ct.Course)), logs[0].Topics[1])
	assert.Equal(t, crypto.Keccak256Hash([]byte(ct.Course)), event.Course)
	assert.Equal(t, id, event.Id)
	assert.Equal(t, ct.Grade, event.Grade)
}

func TestCertificates(t *testing.T) {
	sim, auth, err := simulateBackend()
	if err != nil {
		t.Fatalf("Failed to create simulated backend: %v", err)
	}

	deployedContract, err := deployContract(sim, auth[0])
	if err != nil {
		t.Fatalf("Failed to deploy smart contract: %v", err)
	}

	cert := lib.NewCert()
	instance := cert.Instance(sim.Client(), *deployedContract)

	id := big.NewInt(617)
	ct := lib.CertificatesOutput{
		Name:   "Pine",
		Course: "Reticle",
		Grade:  "A",
		Date:   "11-08-2025",
	}

	_, err = bind.Transact(instance, auth[0], cert.PackIssue(id, ct.Name, ct.Course, ct.Grade, ct.Date))
	if err != nil {
		t.Fatalf("Failed to issue certificate: %v", err)
	}

	sim.Commit()

	out, err := bind.Call(instance, nil, cert.PackCertificates(id), cert.UnpackCertificates)
	if err != nil {
		t.Fatalf("Failed to call Certificates method: %v", err)
	}

	if assert.NotNil(t, out) {
		assert.EqualValues(t, ct, out)
	}
}

func TestModifier(t *testing.T) {
	sim, auth, err := simulateBackend()
	if err != nil {
		t.Fatalf("Failed to create simulated backend: %v", err)
	}

	deployedContract, err := deployContract(sim, auth[0])
	if err != nil {
		t.Fatalf("Failed to deploy smart contract: %v", err)
	}

	cert := lib.NewCert()
	instance := cert.Instance(sim.Client(), *deployedContract)

	id := big.NewInt(250)
	ct := lib.CertificatesOutput{
		Name:   "Pepper",
		Course: "Fury",
		Grade:  "B",
		Date:   "11-08-2025",
	}

	_, err = bind.Transact(instance, auth[1], cert.PackIssue(id, ct.Name, ct.Course, ct.Grade, ct.Date))
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "Access Denied")
}
