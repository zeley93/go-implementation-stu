type Block struct {
	block_hash string
	previous_hash string
	merkle_root string
	nb_transactions int
	transactions[] Transaction
}