package main

import (
	"fmt"

	"github.com/hzhang08/chainnew/blockchain"
)

var blockArray []blockchain.Block

func main() {

	//block := &Block{Index: 0, Timestamp: "565", PreHash: ""}

	//block.Hash = calculateHash(block)

	//newBlock, _ := generateBlock(block)

	//fmt.Println(nowTarget)

	//var hash blockchain.Hash

	fmt.Println(blockchain.GenerateDifficulty())

}
