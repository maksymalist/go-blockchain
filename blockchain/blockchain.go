package blockchain

import (
	"fmt"
)

type BlockChain struct {
	Blocks []*Block
}

func InitBlockChain() *BlockChain {
	fmt.Println("Initializing blockchain... 🚀")
	return &BlockChain{[]*Block{Genesis()}}
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, newBlock)
}
