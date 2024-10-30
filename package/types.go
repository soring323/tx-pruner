package app

type Transaction struct {
	TransactionHash string   `json:"transactionHash"`
	Contracts       []string `json:"contracts"`
}
