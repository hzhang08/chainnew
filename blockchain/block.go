package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"math/big"
	"time"
)

//Block structure
type Block struct {
	Index     int
	Timestamp string
	Hash      string
	PreHash   string
	Nonce     uint64
	Target    *big.Int
	PreBlock  *Block
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

func generateDifficulty() *big.Int {
	difficulty := big.NewInt(1)
	difficulty.Lsh(difficulty, 200)
	return difficulty
}
