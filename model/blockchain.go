package model

import "log"

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

	log.Println("Blockchain :: ", result)
	log.Println("[Start] GetBlockChain")
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
