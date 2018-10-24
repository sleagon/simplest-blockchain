package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// Bill 订单结构体
type Bill struct {
	Amount    int64  `json:"amount"`
	Timestamp int64  `json:"timestamp"`
	From      string `json:"from"`
	To        string `json:"to"`
}

func genBillStr(from, to string, amount int64) (string, error) {
	b := Bill{amount, time.Now().Unix(), from, to}
	bBuffer, err := json.Marshal(b)
	if err != nil {
		return "", err
	}
	return string(bBuffer), nil
}

func main() {
	bc := BlockChain{}
	// 初始化
	bc.Init()

	// 生成一笔订单
	for k := 0; k < 3; k++ {
		b, err := genBillStr("Jimmy", "Tommy", int64(10000*(k+1)))
		if err != nil {
			panic(fmt.Sprintf("Failed to generate new bill of %s", b))
		}
		// 插入到区块链中
		bc.MineBlock(b)
	}
	bc.Print()
}
