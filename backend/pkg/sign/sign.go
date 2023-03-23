package sign

import (
	"context"

	"github.com/lukaracki/flow-go-multisign/backend/pkg/config"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/crypto/cloudkms"
)

// SignVoucher is a function that signs the transaction and returns the signature
func SignVoucher(transaction *flow.Transaction) ([]byte, error) {

	ctx := context.Background()

	// Create a key from the resource ID
	accountKMSKey, err := cloudkms.KeyFromResourceID(config.GCPKmsResourceName)
	if err != nil {
		return nil, err
	}

	// Create a client
	kmsClient, err := cloudkms.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	// Create a signer
	signer, err := kmsClient.SignerForKey(
		ctx,
		accountKMSKey,
	)
	if err != nil {
		return nil, err
	}

	// Sign the envelope
	err = transaction.SignEnvelope(flow.HexToAddress(config.AdminAddress), config.AdminKeyIndex, signer)

	if err != nil {
		return nil, err
	}

	// Return the encoded message
	return transaction.EnvelopeSignatures[len(transaction.EnvelopeSignatures)-1].Signature, nil
}
