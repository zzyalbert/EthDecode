package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/INFURA/go-ethlibs/eth"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("raw tx not given", "use ./eth_decode [raw tx hex]")
		return
	}

	tx := eth.Transaction{}
	err := tx.FromRaw(os.Args[1])
	if err != nil {
		panic(err)
	}

	chainId := (tx.V.UInt64() - 35) / 2
	v := tx.V.UInt64() - 35 - chainId*2

	print := &struct {
		Hash     eth.Data32
		ChainId  uint64
		From     eth.Address
		Nonce    uint64
		To       eth.Address
		Value    uint64
		Data     eth.Data
		Gas      uint64
		GasPrice uint64
		R        eth.Quantity
		S        eth.Quantity
		V        uint64
	}{
		Hash:     tx.Hash,
		ChainId:  chainId,
		From:     tx.From,
		Nonce:    tx.Nonce.UInt64(),
		To:       *tx.To,
		Value:    tx.Value.UInt64(),
		Data:     tx.Input,
		Gas:      tx.Gas.UInt64(),
		GasPrice: tx.GasPrice.UInt64(),
		R:        tx.R,
		S:        tx.S,
		V:        v,
	}

	b, err := json.MarshalIndent(print, "  ", "")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
