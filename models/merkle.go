package models
import (
	"crypto/sha256"
	"encoding/hex"
)
type PseudoMerkle struct {
	transactions [] Transaction
}
func NewPseudoMerkle(transactions [] Transaction) *PseudoMerkle {
	p := new(PseudoMerkle)
	p.transactions = transactions
	return p
}
func (m *PseudoMerkle) HashTransactions(ts [] string) string {
	hash:=""
	for _, tr:= range ts {
		hash = hash + tr
	}
	pseudoMerkleHash := sha256.New()
	pseudoMerkleHash.Write([]byte(hash))
	transactionsHash := hex.EncodeToString(pseudoMerkleHash.Sum(nil))
	return transactionsHash
}