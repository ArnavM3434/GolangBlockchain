package main

import (
	"fmt"

	"github.com/golangblockchain/blockchaingo/blockchain"
)

func main() {
	//fmt.Println(quote.Hello())
	chain := blockchain.InitBlockChain()
	chain.AddBlock("First Block After Genesis")
	chain.AddBlock("Second Block After Genesis")
	chain.AddBlock("Third Block After Genesis")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}

}
