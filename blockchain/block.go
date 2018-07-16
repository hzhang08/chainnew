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

//CalculateHash calculates hash for a block
func CalculateHash(block *Block) string {

	data := string(block.Index) + block.Timestamp + block.PreHash

	h := sha256.New()
	h.Write([]byte(data))
	hashed := h.Sum(nil)

	return hex.EncodeToString(hashed)
}

//GenerateBlock generates a normal block
func GenerateBlock(oldBlock *Block) (*Block, error) {
	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.PreHash = oldBlock.Hash
	newBlock.Hash = CalculateHash(&newBlock)

	return &newBlock, nil
}

//GenerateGenesisBlock generates first block
func GenerateGenesisBlock() (*Block, error) {
	var newBlock Block

	t := time.Now()

	newBlock.Index = 0
	newBlock.Timestamp = t.String()
	newBlock.PreHash = ""
	newBlock.Hash = CalculateHash(&newBlock)

	return &newBlock, nil
}

//GenerateDifficulty generates block difficulty
func GenerateDifficulty() *big.Int {
	difficulty := big.NewInt(1)
	difficulty.Lsh(difficulty, 200)
	return difficulty
}
