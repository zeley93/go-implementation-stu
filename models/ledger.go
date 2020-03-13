package models

type Ledger struct {
	blocks []Block
}

func newLedger() Ledger {
	return Ledger{[]Block{createGenesisBlock()}}
}

func isValid(prevBlock Block, newBlock Block) bool {

	if prevBlock.index+1 != newBlock.index {
		return false
	}
	if prevBlock.blockHash != newBlock.previousHash {
		return false
	}
	if newBlock.genBlockHash() != newBlock.blockHash {
		return false
	}
	return true
}

func getLongestChain(ledger1 *Ledger, ledger2 *Ledger) {
	if len(ledger1) > len(ledger2) {
		return ledger1
	} else if len(ledger1) == len(ledger2) {
		fmt.Println("mork work to do")
	} else {
		return ledger2
	}
}

func (ledger *Ledger) addBlock(merkleRoot string, transactions []Transaction) {
	prevBlock := ledger.getprevBlock()
	newBlock := NewBlock(prevBlock, merkleRoot, transactions)
	if isValid(prevBlock, newBlock) {
		ledger.blocks = append(ledger.blocks, newBlock)
	}
}

func (ledger *Ledger) getprevBlock() Block {
	return ledger.blocks[len(ledger.blocks)]
}
