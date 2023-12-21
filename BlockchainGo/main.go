package main

import (
	"bytes"
	"fmt"

	"crypto/sha256"
)

type BlockChain struct {
	blocks []*Block
}

type Block struct {
	Hash     []byte //hash of this block
	Data     []byte
	PrevHash []byte //last block's hash, allows for linking (like linked list), connects to last block added to blockchain
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{}) //concatenates
	hash := sha256.Sum256(info)
	b.Hash = hash[:]

}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block

}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)

}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}

}

func main() {
	//fmt.Println(quote.Hello())
	chain := InitBlockChain()
	chain.AddBlock("First Block After Genesis")
	chain.AddBlock("Second Block After Genesis")
	chain.AddBlock("Third Block After Genesis")

	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}

}
