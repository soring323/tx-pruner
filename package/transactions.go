package app

// Prune and expand the transactions based on the known contracts
func pruneTransactions(transactions []Transaction, initialContracts []string) []string {
	// Known contracts are tracked and updated dynamically
	knownContracts := initialContracts
	var replayList []string
	// Iterate over the transactions in reverse order
	for i := len(transactions) - 1; i >= 0; i-- {
		tx := transactions[i]
		if containsMatchingContracts(tx.Contracts, knownContracts) {
			replayList = append(replayList, tx.TransactionHash)
			knownContracts = mergeContracts(tx.Contracts, knownContracts)
		}
	}
	return replayList
}

// Helper function to check if the transaction contracts match any known contracts
func containsMatchingContracts(txContracts []string, knownContracts []string) bool {
	knownContractsMap := make(map[string]struct{}, len(knownContracts))
	for _, contract := range knownContracts {
		knownContractsMap[contract] = struct{}{}
	}
	for _, contract := range txContracts {
		if _, exists := knownContractsMap[contract]; exists {
			return true
		}
	}
	return false
}

// Merge the transaction contracts with the known contracts
func mergeContracts(txContracts []string, knownContracts []string) []string {
	knownContractsMap := make(map[string]struct{}, len(knownContracts))
	for _, contract := range knownContracts {
		knownContractsMap[contract] = struct{}{}
	}
	for _, contract := range txContracts {
		if _, exists := knownContractsMap[contract]; !exists {
			knownContractsMap[contract] = struct{}{}
			knownContracts = append(knownContracts, contract)
		}
	}
	return knownContracts
}
