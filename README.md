# Flow - Go - Multisig

This repository is a simple implementation of a multisig service backend and demo frontend, which allows you to execute multisign transactions on the Flow blockchain signed with Google Cloud Platform KMS.

## Prerequisites

Before running the service, make sure you have the following installed:

- Go (at least version 1.18)
- GCP KMS credentials
- Node.js (at least version 14.17.0)

## Installation

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/lukaracki/flow-go-multisign
   ```

2. Cd into the project directory:

   ```bash
   cd flow-go-multisign
   ```

3. Install backend dependencies:

   ```bash
    cd backend && go mod download && cd ..
   ```

4. Install frontend dependencies:

   ```bash
    cd frontend && npm install && cd ..
   ```

5. Setup environment files for both repositoies. Details can be found in the README.md files in each directory.

## Usage

1. Open a new terminal and start the backend service using the following command:

   ```bash
   cd backend && go run cmd/server/main.go
   ```

2. Open a new terminal and start the frontend service using the following command:

   ```bash
   cd frontend && npm run start
   ```
