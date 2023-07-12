package controller

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"main/model"
	"strconv"
	"time"
)

type Block struct {
	PrevHash  string
	Position  int
	Data      model.BookCheckout
	TimeStamp string
	Hash      string
}

func (b *Block) generateHash() {
	bytes, _ := json.Marshal(b.Data)
	data := b.PrevHash + strconv.Itoa(b.Position) + string(bytes) + b.TimeStamp
	hash := sha256.New()
	hash.Write([]byte(data))
	b.Hash = hex.EncodeToString(hash.Sum(nil))
}

func (b *Block) validateHash(hash string) bool {
	b.generateHash()
	return b.Hash == hash
}

func NewBlock(prevBlock *Block, checkoutItem model.BookCheckout) *Block {
	block := &Block{}

	block.PrevHash = prevBlock.Hash
	block.Position = prevBlock.Position + 1
	block.Data = checkoutItem
	block.TimeStamp = time.Now().String()
	block.generateHash()

	return block
}

func GenesisBlock() *Block {
	return NewBlock(&Block{}, model.BookCheckout{IsGenesis: true})
}

type Blockchain struct {
	Blocks []*Block
}

func (bc *Blockchain) AddBlock(data model.BookCheckout) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	block := NewBlock(prevBlock, data)

	if validBlock(block, prevBlock) {
		bc.Blocks = append(bc.Blocks, block)
	}
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{GenesisBlock()}}
}

func validBlock(block, prevBlock *Block) bool {
	if prevBlock.Hash != block.PrevHash {
		return false
	}
	if !block.validateHash(block.Hash) {
		return false
	}
	if prevBlock.Position+1 != block.Position {
		return false
	}

	return true
}

var BlockChain = NewBlockchain()
