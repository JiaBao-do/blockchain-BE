package model

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Block struct {
	Position  int    `json:"position"`
	Data      Data   `json:"data"`
	Timestamp string `json:"timestamp"`
	Hash      string `json:"hash"`
	PrevHash  string `json:"prev_hash"`
}

// generate Hash using given position, timestamp, data, and prevHash
func (b *Block) generateHash() {
	log.Println("[Start] generateHash")
	// get string val of the Data
	bytes, _ := json.Marshal(b.Data)
	// concatenate the dataset
	data := fmt.Sprint(b.Position) + b.Timestamp + string(bytes) + b.PrevHash
	hash := sha256.New()
	hash.Write([]byte(data))
	b.Hash = hex.EncodeToString(hash.Sum(nil))
	log.Println("[End] generateHash")

}

// create new block with given prev block and data
func createBlock(prevBlock *Block, data Data) *Block {
	log.Println("[Start] createBlock")
	block := &Block{}
	block.Position = prevBlock.Position + 1
	block.Timestamp = fmt.Sprint(time.Now().Unix())
	block.Data = data
	block.PrevHash = prevBlock.Hash
	block.generateHash()

	log.Println("Block :: ", block)
	log.Println("[End] createBlock")

	return block
}

func validBlock(block, prevBlock *Block) bool {
	log.Println("[Start] validBlock")
	if prevBlock.Hash != block.PrevHash {
		return false
	}

	if !block.validateHash(block.Hash) {
		return false
	}

	if prevBlock.Position+1 != block.Position {
		return false
	}

	log.Println("[End] validBlock")
	return true
}

// serve as a method to validate
func (b *Block) validateHash(hash string) bool {
	log.Println("[Start] validateHash")
	b.generateHash()
	log.Println("[End] validateHash")
	return b.Hash == hash
}
