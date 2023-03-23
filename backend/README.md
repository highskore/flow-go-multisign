# Flow Multisign Demo - Go Cosign Service

This is a simple service that allows you execute multisignature transactions on the Flow blockchain.

## Prerequisites

Before running the service, make sure you have the following installed:

Go (at least version 1.18)
GCP KMS credentials

## Installation

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/yourusername/flow-go-cosign-service.git
   ```

2. Install the dependencies:

   ```bash
    go mod download
   ```

3. Create a `.env` file in the root directory of the project and add the following environment variables:

   ```bash
    # FLOW
    ADMIN_ADDRESS=0x01cf0e2f2f715450 - Your flow account admin address
    ADMIN_KEY_INDEX=0 - Your flow account admin key index
    # GCP KMS
    GCP_KMS_RESOURCE_NAME='projects/your-project-id/locations/global/keyRings/flow/cryptoKeys/flow-minter-key/cryptoKeyVersions/1' - Your GCP KMS resource name
   ```

## Usage

1. Start the service using the following command:

   ```bash
   go run cmd/server/main.go
   ```

2. The service should now be running and can be accessed via http://localhost:8080.

## Project Structure

    .
    ├── README.md
    ├── api
    │   └── cosign
    │       ├── check.go # Check if the transaction is valid
    │       ├── handler.Go # Handler for the api request
    │       └── process.go # Process signable message into a transaction object
    ├── cmd
    │   └── server
    │       └── main.go # Main package
    ├── go.mod
    ├── go.sum
    ├── internal
    │   └── app
    │       └── server.go # Server implementation
    └── pkg
        ├── config
        │   └── config.go # Configuration files
        ├── interfaces
        │   └── signable.go # Interface for signable object fetched from the client
        └── sign
            └── sign.go # Sign a transaction envelope using GCP KMS
