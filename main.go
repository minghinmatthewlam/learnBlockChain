package main

import (
	"fmt"
	"strconv"

	"github.com/minghinmatthewlam/learnBlockChain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("first block")
	chain.AddBlock("second block")
	chain.AddBlock("third block")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous hash %x\n", block.PrevHash)
		fmt.Printf("data: %s\n", block.Data)
		fmt.Printf("hash: %x\n", block.Hash)

		pow := blockchain.NewProofOfWork(block)
		fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}

}
