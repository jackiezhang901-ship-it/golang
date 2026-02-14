package main

import (
	"bytes"
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
	// 1️⃣ 连接本地 Hardhat 节点
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	// 2️⃣ 私钥（Hardhat node 第一个账户，去掉 0x 前缀）
	privateKeyHex := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	crypto.PubkeyToAddress(*publicKeyECDSA)

	// 3️⃣ 创建 auth（用于发送交易）
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(31337))
	if err != nil {
		log.Fatal(err)
	}

	// 4️⃣ 读取 ABI 文件
	file, err := os.ReadFile(`../solidity/artifacts/contracts/Counter.sol/Counter.json`)
	if err != nil {
		log.Fatal(err)
	}

	var artifact map[string]interface{}
	if err := json.Unmarshal(file, &artifact); err != nil {
		log.Fatal(err)
	}

	abiBytes, _ := json.Marshal(artifact["abi"])
	parsedABI, err := abi.JSON(bytes.NewReader(abiBytes))
	if err != nil {
		log.Fatal(err)
	}

	// 5️⃣ 已部署合约地址
	contractAddress := common.HexToAddress("0x850EC3780CeDfdb116E38B009d0bf7a1ef1b8b38") // 替换成部署时打印的地址

	// 6️⃣ 创建合约实例
	contract := bind.NewBoundContract(contractAddress, parsedABI, client, client, client)

	// =========================
	// 调用 view 函数
	// =========================
	var result []interface{}

	err = contract.Call(&bind.CallOpts{}, &result, "get")
	if err != nil {
		log.Fatal(err)
	}

	count := result[0].(*big.Int)

	fmt.Println("Current count:", count.String())

	// =========================
	// 调用修改状态函数（increment）
	// =========================
	tx, err := contract.Transact(auth, "increment")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Increment tx hash:", tx.Hash().Hex())
}
