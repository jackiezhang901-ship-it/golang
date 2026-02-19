package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// CountIncremented 事件 ABI
const CountABI = `[{"anonymous":false,"inputs":[{"indexed":false,"internalType":"uint256","name":"newCount","type":"uint256"}],"name":"CountIncremented","type":"event"}]`

func main() {
	// 1️⃣ 连接本地 Hardhat WS 节点
	client, err := ethclient.Dial("ws://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// 2️⃣ 合约地址
	contractAddress := common.HexToAddress("0x850EC3780CeDfdb116E38B009d0bf7a1ef1b8b38") // 替换成你的合约地址

	// 3️⃣ 解析 ABI
	parsedABI, err := abi.JSON(strings.NewReader(CountABI))
	if err != nil {
		log.Fatal(err)
	}

	// 4️⃣ 构造实时订阅 FilterQuery
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
		Topics:    [][]common.Hash{{parsedABI.Events["CountIncremented"].ID}},
	}

	// 5️⃣ 创建日志通道并订阅
	logCh := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logCh)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("✅ Subscribed to CountIncremented events... Waiting for events...")

	// 6️⃣ 循环监听事件
	for {
		select {
		case err := <-sub.Err():
			log.Println("subscription error:", err)
		case vLog := <-logCh:
			handleCountEvent(parsedABI, vLog)
		}
	}
}

// 解析事件
func handleCountEvent(parsedABI abi.ABI, vLog types.Log) {
	type CountEvent struct {
		NewCount *big.Int
	}
	var event CountEvent

	// 解析非 indexed 参数
	err := parsedABI.UnpackIntoInterface(&event, "CountIncremented", vLog.Data)
	if err != nil {
		log.Println("Unpack error:", err)
		return
	}

	fmt.Println("🎉 CountIncremented event received! newCount =", event.NewCount.String())
}
