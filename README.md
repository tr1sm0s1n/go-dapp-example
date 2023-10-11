# Go-DApp-Example

DApp example of deploying and interacting with Solidity smart contract using Go.

## 🛠 Built With

![Go](https://img.shields.io/badge/go-2F4F4F?style=for-the-badge&logo=go)
![Gin](https://img.shields.io/badge/gin-2F4F4F?style=for-the-badge&logo=go)
![Solidity](https://img.shields.io/badge/solidity-2F4F4F?style=for-the-badge&logo=solidity)

## ⚙️ Run Locally

Clone the project

```bash
git clone https://github.com/DEMYSTIF/go-dapp-example.git
cd go-dapp-example
```

Install **abigen**

```bash
go install github.com/ethereum/go-ethereum/cmd/abigen@latest
```

Generate Go binding for contract

```bash
abigen --bin bin/Cert.bin --abi bin/Cert.abi --pkg lib --type Cert --out lib/Cert.go
```

Run a blockchain simulation (ganache/geth/foundry), on port **8545** (both http & ws). Copy a valid private key (omit '0x') and paste it in **middlewares.go** (15, 43).

Deploy the contract

```bash
go run ./scripts/deploy.go
```

Copy the contract address from the terminal and paste it in **main.go** (30, 18) and **events.go** (20, 45).

Start the application

```bash
go run .
```

Start the event listener (new terminal)

```bash
go run ./events/events.go
```

Issue a certificate (new terminal)

```bash
curl -X POST http://localhost:8080/issue -H "Content-Type: application/json" -d '{"id": "test-01", "name": "Shepard", "course": "ETH-GO", "grade": "A", "date": "22-11-23"}'
```

Fetch a certificate

```bash
curl http://localhost:8080/fetch/test-01
```

Get event logs

```bash
curl http://localhost:8080/events
```
