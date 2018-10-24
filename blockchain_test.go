package main

import "testing"

// 这里实际测试了Validate和MineBlock，就不再拆分测试了
func TestValidate(t *testing.T) {
	bc := BlockChain{}
	// 初始化
	bc.Init()

	bc.MineBlock("hello block 1")

	if bc.Validate() != true {
		t.Error("Failed to validate right blockchain")
	}

	bc.chain[1].hash = "invalidhash"

	if bc.Validate() != false {
		t.Error("Failed to validate broken blockchain")
	}
}
