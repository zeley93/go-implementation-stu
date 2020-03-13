package models

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

type Block struct {
	index          int
	blockHash      string
	timeStamp      string
	previousHash   string
	merkleRoot     string
	nbTransactions int
	transactions[] Transaction
}

func NewBlock(prevBlock Block, merkleRoot string, transactions[] Transaction) Block{
	b := new(Block)
	b.index = prevBlock.index + 1
	b.previousHash = prevBlock.blockHash
	b.merkleRoot = merkleRoot
	b.nbTransactions = len(transactions)
	b.transactions = transactions
	b.timeStamp = time.Now().String()
	b.blockHash = b.genBlockHash()
	return *b
}

func (block *Block)genBlockHash() string{
	blockString := block.previousHash + block.merkleRoot + block.timeStamp + strconv.Itoa(block.index)
	hash := sha256.New()
	hash.Write([]byte(blockString))
	blockHash := hex.EncodeToString(hash.Sum(nil))
	return blockHash
}

func createGenesisBlock() Block{
	b := new(Block)
	b.index = 0
	b.previousHash = ""
	b.merkleRoot = ""
	b.nbTransactions = 0
	b.timeStamp = time.Now().String()
	b.blockHash = b.genBlockHash()
	return *b
}