package cosign

import (
	"regexp"

	"github.com/onflow/flow-go-sdk"
)

// checkTransaction is a function that checks if all of the relevant transaction fields are valid
// and returns true if it is valid and false if it is not valid
func checkTransaction(transaction *flow.Transaction) bool {
	// Check if the transaction arguments lenght is 0
	return checkArguments(transaction.Arguments) && checkScript(transaction.Script)
}

func checkArguments(arguments [][]byte) bool {
	// If the transaction arguments lenght is 0, return true
	return len(arguments) == 0
}

func checkScript(clientScript []byte) bool {
	pattern := regexp.MustCompile(`\s`)
	// Replace all the whitespace in the client script and compare it to the server script
	return pattern.ReplaceAllString(string(clientScript), "") == pattern.ReplaceAllString(serverScript, "")
}

// Constant variable which holds the server script
const serverScript = `
transaction() {
	prepare(frontendUser: AuthAccount, backendAdmin: AuthAccount) {
	}
}`
