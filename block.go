package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

// Block 区块结构体
type Block struct {
	difficulty int    // 难度系数，越大越难找到需要的nonce，用于控制工作量大小
	nonce      int64  // 没有实在意义，猜出来一个合适的nonce满足difficulty的要求就行了
	index      int64  // 区块序号
	timestamp  int64  // 时间戳
	data       string // 存入区块链的数据，可以是交易或者其他需要存储的信息
	preHash    string // 前一个区块的hash值
	hash       string // 本区块的hash值
}

// Init 初始化区块
func (b *Block) Init(index int64, payload string, preHash string, difficulty int) {
	b.data = payload
	b.index = index
	b.preHash = preHash
	b.difficulty = difficulty
	b.timestamp = time.Now().UnixNano()
	// 这里初始化成0，实际可以初始化成任何值
	b.nonce = 0
}

// CalculateHash 计算区块hash值
func (b *Block) CalculateHash() string {
	// 具体计算方式也不一定，这里是用最简单的方式拼接一个字符串，算SHA256值
	input := fmt.Sprintf("%d:%d:%s:%s", b.nonce, b.index, b.data, b.preHash)
	return fmt.Sprintf("%x", sha256.Sum256([]byte(input)))
}

// FixHash 修复hash值，传入空字符串会自动计算一个
func (b *Block) FixHash(hash string) {
	if hash != "" {
		b.hash = hash
		return
	}
	b.hash = b.CalculateHash()
}
