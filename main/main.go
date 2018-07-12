package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"
)

//Block structure
type Block struct {
	Index     int
	Timestamp string
	Hash      string
	PreHash   string
	Nonce     int
	Target    big.Int
}

var blockchain []Block

const nowTarget = 1 << 100

func main() {

	block := &Block{Index: 0, Timestamp: "565", PreHash: ""}

	block.Hash = calculateHash(block)

	newBlock, _ := generateBlock(block)

	fmt.Println(nowTarget)
}

func calculateHash(block *Block) string {

	data := string(block.Index) + block.Timestamp + block.PreHash

	h := sha256.New()
	h.Write([]byte(data))
	hashed := h.Sum(nil)

	return hex.EncodeToString(hashed)
}

func generateBlock(oldBlock *Block) (*Block, error) {
	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.PreHash = oldBlock.Hash
	newBlock.Hash = calculateHash(&newBlock)

	return &newBlock, nil
}
