package app

import (
	"fmt"
)

func PruneTransactions() {
	// Initial contracts to start the process
	initialContracts := []string{"0xUniswap"}

	// Sample transactions (can be replaced with actual blockchain data)
	transactions := []Transaction{
		{TransactionHash: "0xTx0", Contracts: []string{"0xWBTC"}},
		{TransactionHash: "0xTx1", Contracts: []string{"0xUniswap", "0xWETH"}},
		{TransactionHash: "0xTx2", Contracts: []string{"0xUSDC", "0xUSDT"}},
		{TransactionHash: "0xTx3", Contracts: []string{"0xNFT", "0xWETH"}},
		{TransactionHash: "0xTx4", Contracts: []string{"0xNFT", "0xWETH"}},
	}

	// Execute the prune function to get the replay list
	replayList := pruneTransactions(transactions, initialContracts)

	// Output the list of transactions to replay
	fmt.Println("Transactions to replay:", replayList)
}
