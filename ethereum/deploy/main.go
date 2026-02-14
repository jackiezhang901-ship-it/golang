package main

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	// 连接本地 Hardhat 节点
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	// 读取私钥（Hardhat 默认账户）
	privateKey, err := crypto.HexToECDSA("de9be858da4a475276426320d5e9262ecfc3ba460bfac56360bfa6c4c28b4ee0")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, _ := client.PendingNonceAt(context.Background(), fromAddress)

	gasPrice, _ := client.SuggestGasPrice(context.Background())

	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(31337))

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = gasPrice

	// 读取 Hardhat 编译结果
	file, _ := os.ReadFile("../solidity/artifacts/contracts/Counter.sol/Counter.json")

	var result map[string]interface{}
	json.Unmarshal(file, &result)

	abiBytes, _ := json.Marshal(result["abi"])
	bytecode := common.FromHex(result["bytecode"].(string))

	parsedABI, _ := abi.JSON(bytes.NewReader(abiBytes))

	// 部署
	address, tx, _, err := bind.DeployContract(auth, parsedABI, bytecode, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deploy tx:", tx.Hash().Hex())
	fmt.Println("Contract address:", address.Hex())
}
