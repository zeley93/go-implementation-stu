package models

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)
type Transaction struct {
	timestamp string 
	transaction_id string
	sender_address string
	amount int
	receiver_address string
}

func (t *Transaction) GetTransactionId() string {
	return t.transaction_id
}

func NewTransaction(senderAddress string, amount int, receiverAddress string) Transaction {
	t:=new(Transaction)
	t.timestamp = string(time.Now().Unix())
	t.sender_address = senderAddress
	t.amount = amount
	t.receiver_address = receiverAddress
	t.transaction_id = t.hashTransaction()
	return *t
}

func (t *Transaction)hashTransaction() string {
	concat := t.sender_address + string(t.amount) + t.receiver_address +  t.timestamp
	hash := sha256.New()
	hash.Write([]byte(concat))
	transactionHash := hex.EncodeToString(hash.Sum(nil))
	return transactionHash
}
