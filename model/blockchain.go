package model

import (
	"encoding/json"
	"log"
)

type Blockchain struct {
	blocks []*Block
}

var BlockChain *Blockchain

// Initialize blockchain
func Initialize() {
	BlockChain = newBlockchain()
}
func newBlockchain() *Blockchain {
	return &Blockchain{[]*Block{genesisBlock()}}
}
func genesisBlock() *Block {

	return createBlock(&Block{}, Data{IsGenesis: true})
}

/*
* method to get all block chain
 */
func GetBlockChain() []Block {

	log.Println("[Start] GetBlockChain")
	if BlockChain == nil {
		log.Println("Blockchain :: ", nil)
		log.Println("[End] GetBlockChain")
		return []Block{}
	}

	result := make([]Block, len(BlockChain.blocks))

	for i, block := range BlockChain.blocks {
		result[i] = *block

	}

	/*  {
		position:
		data: {
			content:
			is_genesis:
		}
		timestamp:
		hash:
		preve_hash:
	}
	*/

	val, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		log.Println("[End] GetBlockChain")
		return nil
	}
	log.Println("Blockchain :: ", string(val))
	log.Println("[End] GetBlockChain")
	return result
}

// to serve as entry point for public to add new block to blockchain
func PrepareBlock(data Data) {
	BlockChain.addBlock(data)
}

func (bc *Blockchain) addBlock(data Data) {
	log.Print("[Start] addBlock")

	prevBlock := bc.blocks[len(bc.blocks)-1]

	block := createBlock(prevBlock, data)

	if validBlock(block, prevBlock) {
		bc.blocks = append(bc.blocks, block)
	}
	GetBlockChain()
	log.Print("[End] addBlock")

}
