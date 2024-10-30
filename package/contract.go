package app

// Helper function to check if a contract exists in the array
func containsContract(contract string, contracts []string) bool {
	for _, c := range contracts {
		if c == contract {
			return true
		}
	}
	return false
}

// Add new contracts from a transaction that are not yet in the contract list
func addNewContracts(txContracts []string, knownContracts *[]string) bool {
	updated := false
	for _, contract := range txContracts {
		if !containsContract(contract, *knownContracts) {
			*knownContracts = append(*knownContracts, contract)
			updated = true
		}
	}
	return updated
}
