package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

type BlockChain struct {
	blocks []*Block
}

// join previous blocks with new blocks
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
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
	chain := InitBlockChain()
	chain.AddBlock("First BLOCK after GENESIS")
	chain.AddBlock("Second BLOCK after GENESIS")
	chain.AddBlock("Third BLOCK after GENESIS")

	for _, block := range chain.blocks {
		fmt.Printf("Previous hash %x\n", block.PrevHash)
		fmt.Printf("Data %s\n", block.Data)
		fmt.Printf("Hash %x\n", block.Hash)
	}
}
