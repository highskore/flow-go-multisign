package cosign

import (
	"encoding/hex"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lukaracki/flow-go-multisign/backend/pkg/interfaces"
	"github.com/lukaracki/flow-go-multisign/backend/pkg/sign"
	"github.com/onflow/flow-go-sdk"
)

// ProcessSignable is a function that processes the signable object, first it decodes the hex string
// and then it decodes the transaction. It then checks the transaction and if it is valid it will
// sign the transaction and return the signature.
func ProcessSignable(c *gin.Context, signable *interfaces.Signable) ([]byte, error) {
	decoded, err := hex.DecodeString(signable.Message[64:])

	if err != nil {
		return nil, fmt.Errorf("error decoding hex string: %w", err)
	}

	transaction, err := flow.DecodeTransaction(decoded)

	if err != nil {
		return nil, fmt.Errorf("error decoding transaction: %w", err)
	}

	if checkTransaction(transaction) {
		return sign.SignVoucher(transaction)
	} else {
		err = fmt.Errorf("invalid request: you are not authorized to request this signature")
		return nil, err
	}
}
