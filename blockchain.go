package main

import (
	"fmt"
	"strings"
)

// BlockChain 区块链结构体
type BlockChain struct {
	chain []Block
}

// Init 初始化区块链结构体
func (bc *BlockChain) Init() {
	var b Block
	// 这里随便写了一段信息作为创世块，比如比特币取的是那一天的报纸标题
	b.Init(0, "This is this genesis block!", "", 2)
	b.FixHash("")
	bc.chain = []Block{b}
}

// GetLastlock 获取最后一个区块
func (bc BlockChain) GetLastlock() *Block {
	return &bc.chain[len(bc.chain)-1]
}

// MineBlock 挖取一个最新的区块
func (bc *BlockChain) MineBlock(payload string) {
	l := bc.GetLastlock()
	var nb Block
	nb.Init(l.index+1, payload, l.hash, l.difficulty)
	var hash string
	for {
		hash = nb.CalculateHash()
		if hash[:nb.difficulty] == strings.Repeat("0", nb.difficulty) {
			break
		}
		nb.nonce++
	}
	nb.FixHash(hash)
	bc.chain = append(bc.chain, nb)
}

// Validate 判断区块链是否合法，是否被篡改
func (bc BlockChain) Validate() bool {
	l := len(bc.chain)
	for k := 0; k < l; k++ {
		if bc.chain[k].CalculateHash() != bc.chain[k].hash {
			return false
		}
		if k == 0 {
			continue
		}
		if bc.chain[k].preHash != bc.chain[k-1].hash {
			return false
		}
	}
	return true
}

// Print 打印区块链信息
func (bc BlockChain) Print() {
	fmt.Printf("Is this blockchain validate? %t\n", bc.Validate())
	for _, v := range bc.chain {
		fmt.Println("-----------------------------------------------")
		fmt.Printf("%+v\n", v)
	}
}
