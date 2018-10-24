package main

import (
	"testing"
)

func TestCalculateHash(t *testing.T) {
	b := Block{}
	b.Init(3, `{"amount":30000,"timestamp":1540385923,"from":"Jimmy","to":"Tommy"}`, "00caad8bb0ca353c67b956f4a9ba8a683f22c98a405bbc18dc772e7c0cb5d1d8", 2)
	b.nonce = 28
	if b.CalculateHash() != "00295a92953b02198714adfc0df92b82868fc27fa09995de0da446e819916b44" {
		t.Error("Wrong hash got")
	}
}

func TestFixHash(t *testing.T) {
	b := Block{}
	b.Init(3, `{"amount":30000,"timestamp":1540385923,"from":"Jimmy","to":"Tommy"}`, "00caad8bb0ca353c67b956f4a9ba8a683f22c98a405bbc18dc772e7c0cb5d1d8", 2)
	b.FixHash("")
	if b.hash != "50cb84de5923a6d13c2bd30f3496ad7e4d4d1c58aa3237de847db5685f742bc7" {
		t.Error("Failed to auto fix hash")
	}
	b.Init(3, `{"amount":30000,"timestamp":1540385923,"from":"Jimmy","to":"Tommy"}`, "00caad8bb0ca353c67b956f4a9ba8a683f22c98a405bbc18dc772e7c0cb5d1d8", 2)
	b.FixHash("xxx")
	if b.hash != "xxx" {
		t.Error("Failed to fix hash with particular string")
	}
}
