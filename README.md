# Go-DApp-Example

DApp example of deploying and interacting with Solidity smart contract using Go.

## 🛠 Built With

[![Go Badge](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=fff&style=for-the-badge)](https://go.dev/)
[![Gin Badge](https://img.shields.io/badge/Gin-008ECF?logo=gin&logoColor=fff&style=for-the-badge)](https://gin-gonic.com/)
[![Geth Badge](https://img.shields.io/badge/Geth-3C3C3D?logo=ethereum&logoColor=fff&style=for-the-badge)](https://geth.ethereum.org/)
[![Solidity Badge](https://img.shields.io/badge/Solidity-363636?logo=solidity&logoColor=fff&style=for-the-badge)](https://soliditylang.org/)

## ⚙️ Run Locally

Clone the project.

```bash
git clone https://github.com/tr1sm0s1n/go-dapp-example.git
cd go-dapp-example
```

Install **abigen**.

```bash
go install github.com/ethereum/go-ethereum/cmd/abigen@latest
```

Generate Go binding for contract.

```bash
abigen --bin lib/Cert.bin --abi lib/Cert.abi --pkg lib --type Cert --out lib/Cert.go
```

Run a blockchain simulation on port **8545** (both http & ws). Copy a valid private key (omit '0x') and paste it in **.env**.

```bash
PRIVATE_KEY=<private-key>
```

Deploy the contract.

```bash
go run ./cmd/deploy.go
```

Copy the contract address from the terminal and paste it in **.env**.

```bash
CONTRACT_ADDRESS=<contract-address>
```

Start the application.

```bash
go run .
```

Start the event listener (new terminal).

```bash
go run ./events/events.go
```

Issue a certificate (new terminal).

```bash
curl -X POST http://localhost:8080/issue -H "Content-Type: application/json" -d '{"id": "test-01", "name": "Shepard", "course": "ETH-GO", "grade": "A", "date": "22-11-23"}'
```

Fetch a certificate.

```bash
curl http://localhost:8080/fetch/test-01
```

Get event logs.

```bash
curl http://localhost:8080/events
```
